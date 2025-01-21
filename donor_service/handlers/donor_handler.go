package handlers

import (
	"context"
	"donor-service/dtos"
	"donor-service/models"
	"donor-service/proto/pb"
	"donor-service/repository"
	"donor-service/service"
	"net/http"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type DonorHandler struct {
	DonorRepository  repository.DonorRepository
	PaymentService   service.PaymentService
	FoundationClient pb.FoundationServiceClient
	RestaurantClient pb.RestaurantServiceClient
}

func NewDonorHandlerImpl(donorRepository repository.DonorRepository, paymentService service.PaymentService, foundationClient pb.FoundationServiceClient, restaurantClient pb.RestaurantServiceClient) *DonorHandler {
	return &DonorHandler{
		DonorRepository:  donorRepository,
		PaymentService:   paymentService,
		FoundationClient: foundationClient,
		RestaurantClient: restaurantClient,
	}
}

func (dh *DonorHandler) GetDonorByID(c echo.Context) error {
	donorID := c.Param("id")

	donor, err := dh.DonorRepository.GetDonorByID(donorID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.JSON(http.StatusNotFound, map[string]string{"message": "Donor not found"})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to retrieve donor detail"})
	}

	res := dtos.DonorData{
		ID:      donorID,
		Name:    donor.Name,
		UserID:  donor.UserID,
		Balance: donor.Balance,
	}

	return c.JSON(http.StatusOK, res)
}

func (dh *DonorHandler) TopUp(c echo.Context) error {
	donorID := c.Param("donorID")

	_, err := dh.DonorRepository.GetDonorByID(donorID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.JSON(http.StatusNotFound, map[string]string{"message": "Donor not found"})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to retrieve donor detail"})
	}

	req := new(dtos.TopUpRequest)

	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid input"})
	}

	if req.Amount < 10000 {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Top up amount must be at least Rp. 10.000"})
	}

	newUUID := uuid.New().String()

	invoice, err := dh.PaymentService.CreateInvoice(newUUID, req.Amount)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to create invoice"})
	}

	topUp := &models.TopUp{
		ID:            newUUID,
		DonorID:       donorID,
		Amount:        req.Amount,
		InvoiceID:     invoice.InvoiceID,
		InvoiceUrl:    invoice.InvoiceUrl,
		PaymentMethod: invoice.Method,
		Status:        invoice.Status,
	}

	if err := dh.DonorRepository.TopUp(topUp); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to top up"})
	}

	return c.JSON(http.StatusCreated, topUp)
}

func (dh *DonorHandler) GetTopUpHistory(c echo.Context) error {
	donorID := c.Param("donorID")

	_, err := dh.DonorRepository.GetDonorByID(donorID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.JSON(http.StatusNotFound, map[string]string{"message": "Donor not found"})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to retrieve donor detail"})
	}

	topUps, err := dh.DonorRepository.GetTopUpHistory(donorID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to retrieve top up history"})
	}

	return c.JSON(http.StatusOK, topUps)
}

func (dh *DonorHandler) UpdateTopUpStatus(c echo.Context) error {
	webhookToken := c.Request().Header.Get("x-callback-token")
	if webhookToken != os.Getenv("XENDIT_WEBHOOK_TOKEN") {
		return c.JSON(http.StatusUnauthorized, map[string]string{"message": "Invalid webhook token"})
	}

	req := new(dtos.XenditWebhookRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid input"})
	}

	topUp, err := dh.DonorRepository.UpdateTopUpStatus(req.ExternalID, req.Status, req.PaymentMethod, req.CompletedAt)
	if err != nil {
		if err.Error() == "not found" {
			return c.JSON(http.StatusNotFound, map[string]string{"message": "Top up not found"})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to update top up status"})
	}

	if err := dh.DonorRepository.AddDonorBalance(topUp.DonorID, topUp.Amount); err != nil {
		if err.Error() == "not found" {
			return c.JSON(http.StatusNotFound, map[string]string{"message": "Donor not found"})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to update donor balance"})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Successfully updated top up status"})
}

func (dh *DonorHandler) Donate(c echo.Context) error {
	donorID := c.Param("donorID")

	req := new(dtos.DonateRequest)

	if err := c.Bind(req); err != nil || req.Quantity <= 0 {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid input"})
	}

	// Check if donor exist based on donor id (database)
	donor, err := dh.DonorRepository.GetDonorByID(donorID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.JSON(http.StatusNotFound, map[string]string{"message": "Donor not found"})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to retrieve donor detail"})
	}

	// Check if order exist based on order id (database)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	orders, err := dh.FoundationClient.GetOrderByID(ctx, &pb.OrderID{Id: req.OrderID})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	// Check if order quantity == desired quantity
	if orders.Quantity == orders.DesiredQuantity {
		return c.JSON(http.StatusUnprocessableEntity, map[string]string{"message": "This order has met it's desired quantity"})
	}

	// Check if added orders quantity is bigger than desired quantity
	orders.Quantity += req.Quantity
	if orders.Quantity > orders.DesiredQuantity {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid quantity"})
	}

	// Get meals detail with meals id (database)
	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	meal, err := dh.RestaurantClient.GetMealByID(ctx, &pb.MealID{Id: orders.MealsId})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	// Check if requested quantity is bigger than meal stock
	if req.Quantity > meal.Stock {
		return c.JSON(http.StatusUnprocessableEntity, map[string]string{"message": "Quantity is larger than meal stock"})
	}

	// Calculate total price and check if donor has sufficient balance
	totalPrice := meal.Price * float32(req.Quantity)
	if totalPrice > float32(donor.Balance) {
		return c.JSON(http.StatusUnprocessableEntity, map[string]string{"message": "Insufficient balance"})
	}

	donationID := uuid.New().String()

	// Deduct meals stock and update meals table (database)
	mealRequest := &pb.PrepareDeductMealStockRequest{
		DonationId: donationID,
		MealId:     meal.Id,
		Quantity:   req.Quantity,
	}

	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	response, err := dh.RestaurantClient.PrepareDeductMealStock(ctx, mealRequest)
	if err != nil || !response.Success {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to update meal via gRPC: " + response.GetMessage(),
		})
	}

	// Start transaction
	tx, err := dh.DonorRepository.BeginTransaction()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	defer tx.Rollback()

	// Deduct donor balance
	if err := dh.DonorRepository.DeductDonorBalance(tx, donorID, float64(totalPrice)); err != nil {
		_, _ = dh.RestaurantClient.RollbackDeductMealStock(ctx, &pb.RollbackDeductMealStockRequest{DonationId: donationID})
		if err.Error() == "not found" {
			return c.JSON(http.StatusNotFound, map[string]string{"message": "Donor not found"})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to deduct donor balance"})
	}

	// Update orders table
	order := &pb.PrepareAddOrderQuantityRequest{
		DonationId: donationID,
		OrderId:    orders.Id,
		Quantity:   req.Quantity,
	}

	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	response2, err := dh.FoundationClient.PrepareAddOrderQuantity(ctx, order)
	if err != nil || !response2.Success {
		_, _ = dh.RestaurantClient.RollbackDeductMealStock(ctx, &pb.RollbackDeductMealStockRequest{DonationId: donationID})
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to update order via gRPC: " + response2.GetMessage(),
		})
	}

	donation := &models.Donation{
		ID:       donationID,
		OrderID:  req.OrderID,
		DonorID:  donorID,
		Quantity: req.Quantity,
	}

	// Create donation
	if err := dh.DonorRepository.CreateDonation(tx, donation); err == nil {
		_, _ = dh.RestaurantClient.RollbackDeductMealStock(ctx, &pb.RollbackDeductMealStockRequest{DonationId: donationID})
		_, _ = dh.FoundationClient.RollbackAddOrderQuantity(ctx, &pb.RollbackAddOrderQuantityRequest{DonationId: donationID})
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to donate"})
	}

	if err := tx.Commit().Error; err != nil {
		_, _ = dh.RestaurantClient.RollbackDeductMealStock(ctx, &pb.RollbackDeductMealStockRequest{DonationId: donationID})
		_, _ = dh.FoundationClient.RollbackAddOrderQuantity(ctx, &pb.RollbackAddOrderQuantityRequest{DonationId: donationID})
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to commit donation transaction"})
	}

	_, err = dh.RestaurantClient.CommitDeductMealStock(ctx, &pb.CommitDeductMealStockRequest{DonationId: donationID})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to commit meal via gRPC"})
	}

	_, err = dh.FoundationClient.CommitAddOrderQuantity(ctx, &pb.CommitAddOrderQuantityRequest{DonationId: donationID})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to commit add order quantity via gRPC"})
	}

	return c.JSON(http.StatusOK, donation)
}

func (dh *DonorHandler) GetDonationHistory(c echo.Context) error {
	donorID := c.Param("donorID")

	_, err := dh.DonorRepository.GetDonorByID(donorID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.JSON(http.StatusNotFound, map[string]string{"message": "Top up not found"})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to retrieve top up detail"})
	}

	donations, err := dh.DonorRepository.GetDonationHistory(donorID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to retrieve donation history"})
	}

	return c.JSON(http.StatusOK, donations)
}
