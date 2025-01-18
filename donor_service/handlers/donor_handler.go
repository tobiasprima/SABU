package handlers

import (
	"donor-service/dtos"
	"donor-service/models"
	"donor-service/repository"
	"donor-service/service"
	"net/http"
	"os"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type DonorHandler struct {
	DonorRepository repository.DonorRepository
	PaymentService  service.PaymentService
}

func NewDonorHandlerImpl(donorRepository repository.DonorRepository, paymentService service.PaymentService) *DonorHandler {
	return &DonorHandler{
		DonorRepository: donorRepository,
		PaymentService:  paymentService,
	}
}

func (dh *DonorHandler) GetDonorByID(c echo.Context) error {
	donorID := c.Param("id")

	donor, err := dh.DonorRepository.GetDonorByID(donorID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return echo.NewHTTPError(http.StatusNotFound, map[string]string{"message": "donor not found"})
		}
		return echo.NewHTTPError(http.StatusInternalServerError, map[string]string{"message": "failed to retrieve donor detail"})
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

func (dh *DonorHandler) UpdateTopUpStatus(c echo.Context) error {
	webhookToken := c.Request().Header.Get("x-callback-token")
	if webhookToken != os.Getenv("XENDIT_WEBHOOK_TOKEN") {
		return c.JSON(http.StatusUnauthorized, map[string]string{"message": "Invalid webhook token"})
	}

	req := new(dtos.XenditWebhookRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid input"})
	}

	topUp, err := dh.DonorRepository.GetTopUpByID(req.ExternalID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.JSON(http.StatusNotFound, map[string]string{"message": "Top up not found"})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to retrieve top up detail"})
	}

	if err := dh.DonorRepository.UpdateTopUpStatus(topUp, req.CompletedAt, req.Status, req.PaymentMethod); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to update top up status"})
	}

	if err := dh.DonorRepository.UpdateDonorBalance(topUp.DonorID, topUp.Amount); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to update donor balance"})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Successfully updated top up status"})
}
