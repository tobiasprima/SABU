package handlers

import (
	"foundation-service/dtos"
	"foundation-service/models"
	"foundation-service/repository"

	"foundation-service/utils"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type FoundationHandler struct {
	FoundationRepository repository.FoundationRepository
}

func NewFoundationHandlerImpl(foundationRepository repository.FoundationRepository) *FoundationHandler {
	return &FoundationHandler{FoundationRepository: foundationRepository}
}

// GetFoundationByID godoc
// @Summary      Get foundation by ID
// @Description  Retrieve details of a foundation by its ID
// @Tags         Foundation
// @Param        foundation_id path string true "Foundation ID"
// @Success      200 {object} dtos.FoundationData
// @Failure      400 {object} map[string]string
// @Failure      404 {object} map[string]string
// @Failure      500 {object} map[string]string
// @Router       /foundation/{foundation_id} [get]
func (fh *FoundationHandler) GetFoundationByID(c echo.Context) error {
	foundationID := c.Param("foundation_id")

	foundation, err := fh.FoundationRepository.GetFoundationByID(foundationID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return echo.NewHTTPError(http.StatusNotFound, map[string]string{"message": "foundation not found"})
		}
		return echo.NewHTTPError(http.StatusInternalServerError, map[string]string{"message": "failed to retrieve foundation detail"})
	}

	res := dtos.FoundationData{
		ID:      foundationID,
		Name:    foundation.Name,
		UserID:  foundation.UserID,
		Address: foundation.Address,
	}

	return c.JSON(http.StatusOK, res)
}

// AddOrderlist godoc
// @Summary      Add a new order list
// @Description  Create a new order list for the specified foundation
// @Tags         OrderList
// @Param        foundation_id path string true "Foundation ID"
// @Success      201 {object} map[string]string
// @Failure      400 {object} map[string]string
// @Failure      500 {object} map[string]string
// @Router       /foundation/add-orderlist/{foundation_id} [post]
func (fh *FoundationHandler) AddOrderlist(c echo.Context) error {
	foundationID := c.Param("foundation_id")
	if foundationID == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Foundation ID is required"})
	}

	orderlist := &models.OrderList{
		FoundationID: foundationID,
		Status:       "unpaid",
	}

	if err := fh.FoundationRepository.AddOrderlist(orderlist); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to add orderlist"})
	}

	return c.JSON(http.StatusCreated, map[string]string{"message": "Orderlist created successfully"})
}

// AddOrder godoc
// @Summary      Add orders to an order list
// @Description  Add multiple orders to an existing order list
// @Tags         Order
// @Param        orderlist_id path string true "Order List ID"
// @Param        body body dtos.OrderRequest true "Order request payload"
// @Success      201 {object} map[string]interface{}
// @Failure      400 {object} map[string]string
// @Failure      500 {object} map[string]string
// @Router       /foundation/add-order/{orderlist_id} [post]
func (fh *FoundationHandler) AddOrder(c echo.Context) error {
	orderListID := c.Param("orderlist_id")
	if orderListID == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Orderlist ID is required"})
	}

	req := new(dtos.OrderRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
	}

	var orders []models.Order
	for _, orderReq := range req.Orders {
		order := models.Order{
			OrderListID:     orderListID,
			MealsID:         orderReq.MealsID,
			DesiredQuantity: orderReq.DesiredQuantity,
			Quantity:        0,
		}
		orders = append(orders, order)
	}

	if err := fh.FoundationRepository.AddOrders(orders); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to add orders"})
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"message": "Orders added successfully",
		"orders":  orders,
	})
}

// GetOrder godoc
// @Summary      Get orders by order list ID
// @Description  Retrieve all orders for a specific order list
// @Tags         Order
// @Param        orderlist_id path string true "Order List ID"
// @Success      200 {object} map[string]interface{}
// @Failure      400 {object} map[string]string
// @Failure      404 {object} map[string]string
// @Failure      500 {object} map[string]string
// @Router       /foundation/get-order/{orderlist_id} [get]
func (fh *FoundationHandler) GetOrder(c echo.Context) error {
	// Extract orderlist_id from the URL parameter
	orderlistID := c.Param("orderlist_id")
	if orderlistID == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Orderlist ID is required"})
	}

	// Fetch the OrderList from the database
	var orderlist models.OrderList
	if err := fh.FoundationRepository.GetOrderlistByID(orderlistID, &orderlist); err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.JSON(http.StatusNotFound, map[string]string{"error": "Orderlist not found"})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to fetch orderlist"})
	}

	// Fetch the Orders associated with the OrderList
	var orders []models.Order
	if err := fh.FoundationRepository.GetOrdersByOrderlistID(orderlistID, &orders); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to fetch orders"})
	}

	// Combine the OrderList and Orders into a response
	response := map[string]interface{}{
		"orderlist": orderlist,
		"orders":    orders,
	}

	return c.JSON(http.StatusOK, response)
}

func (fh *FoundationHandler) GetOrderById(c echo.Context) error {
	// Extract order_id from URL parameter
	orderID := c.Param("order_id")
	if orderID == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Order ID is required"})
	}

	// Fetch the order by ID
	order, err := fh.FoundationRepository.GetOrderByID(orderID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.JSON(http.StatusNotFound, map[string]string{"error": "Order not found"})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to fetch order"})
	}

	// Return the order in the response
	return c.JSON(http.StatusOK, order)
}

// CompleteOrder godoc
// @Summary      Mark an order list as complete
// @Description  Mark an order list as complete and send an email notification
// @Tags         OrderList
// @Param        orderlist_id path string true "Order List ID"
// @Param        body body dtos.CompleteOrderRequest true "Complete order request payload"
// @Success      200 {object} map[string]string
// @Failure      400 {object} map[string]string
// @Failure      404 {object} map[string]string
// @Failure      500 {object} map[string]string
// @Router       /foundation/complete-order/{orderlist_id} [post]
func (fh *FoundationHandler) CompleteOrder(c echo.Context) error {
	// Extract orderlist_id from the URL or payload
	orderListID := c.Param("orderlist_id")
	if orderListID == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "OrderList ID is required"})
	}

	var payload dtos.CompleteOrderRequest
	if err := c.Bind(&payload); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid payload"})
	}

	if payload.Email == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Email is required"})
	}

	var orderList models.OrderList
	if err := fh.FoundationRepository.GetOrderlistByID(orderListID, &orderList); err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.JSON(http.StatusNotFound, map[string]string{"error": "Orderlist not found"})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to fetch orderlist"})
	}

	// Fetch all orders for the given orderlist
	orders, err := fh.FoundationRepository.GetOrdersArrayByOrderListID(orderListID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to fetch orders"})
	}

	// Check if all orders meet the desired quantity
	for _, order := range orders {
		if order.Quantity < order.DesiredQuantity {
			return c.JSON(http.StatusOK, map[string]string{
				"message": "OrderList is not complete yet",
			})
		}
	}

	// If all orders are complete, update the OrderList status to "complete"
	if err := fh.FoundationRepository.UpdateOrderListStatus(orderListID, "complete"); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to update OrderList status"})
	}

	foundation, err := fh.FoundationRepository.GetFoundation(orderListID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to fetch foundation details"})
	}

	//Send email notification
	email := payload.Email
	name := foundation.Name
	err = utils.SendCompletionEmail(email, name)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to send email notification"})
	}

	//Return success response
	return c.JSON(http.StatusOK, map[string]string{
		"message": "OrderList marked as complete",
	})
}
