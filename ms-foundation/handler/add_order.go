package handler

import (
	"ms-foundation/config"
	"ms-foundation/entity"
	"net/http"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func AddOrder(c echo.Context) error {
	// Extract user claims from JWT
	user := c.Get("user")
	if user == nil {
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"message": "Unauthorized access",
		})
	}

	claims, ok := user.(jwt.MapClaims)
	if !ok {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Failed to parse user claims",
		})
	}

	userIDFloat, ok := claims["id"].(float64)
	if !ok {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "User ID not found in claims",
		})
	}
	userID := int(userIDFloat)

	// Find the foundation_id associated with the user_id
	var foundation entity.Foundation
	if err := config.DB.Where("user_id = ?", userID).First(&foundation).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"message": "Foundation not found for this user",
			"error":   err.Error(),
		})
	}

	// Parse the incoming request
	var req entity.OrderRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Invalid request payload",
			"error":   err.Error(),
		})
	}

	// Validate required fields
	if err := c.Validate(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Validation error",
			"error":   err.Error(),
		})
	}

	// Check if the orderlist_id exists
	var orderList entity.OrderList
	if err := config.DB.Where("id = ? AND foundation_id = ?", req.OrderListID, foundation.ID).First(&orderList).Error; err != nil {
		if err.Error() == "record not found" {
			// Create a new order list if it doesn't exist
			newOrderList := entity.OrderList{
				ID:           req.OrderListID,
				FoundationID: foundation.ID,
			}
			if err := config.DB.Create(&newOrderList).Error; err != nil {
				return c.JSON(http.StatusInternalServerError, map[string]interface{}{
					"message": "Failed to create new order list",
					"error":   err.Error(),
				})
			}
			orderList = newOrderList
		} else {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"message": "Failed to query order list",
				"error":   err.Error(),
			})
		}
	}

	// Add new orders to the order list
	for _, o := range req.Orders {
		// Check if an existing order already exists with the same meals_id and orderlist_id
		var existingOrder entity.Order
		err := config.DB.Where("orderlist_id = ? AND meals_id = ?", req.OrderListID, o.MealsID).First(&existingOrder).Error

		if err == nil {
			// If order exists, add the new desired quantity to the existing order's quantity
			existingOrder.DesiredQuantity += o.DesiredQuantity
			if err := config.DB.Save(&existingOrder).Error; err != nil {
				return c.JSON(http.StatusInternalServerError, map[string]interface{}{
					"message": "Failed to update existing order",
					"error":   err.Error(),
				})
			}
		} else if err.Error() == "record not found" {
			// If order doesn't exist, create a new order
			newOrder := entity.Order{
				OrderListID:     req.OrderListID,
				MealsID:         o.MealsID,
				DesiredQuantity: o.DesiredQuantity,
			}
			if err := config.DB.Create(&newOrder).Error; err != nil {
				return c.JSON(http.StatusInternalServerError, map[string]interface{}{
					"message": "Failed to create new order",
					"error":   err.Error(),
				})
			}
		}
	}

	// Check if the order list can be marked as completed
	var orders []entity.Order
	if err := config.DB.Where("orderlist_id = ?", req.OrderListID).Find(&orders).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Failed to fetch orders",
			"error":   err.Error(),
		})
	}

	allCompleted := true
	for _, order := range orders {
		if order.Quantity < order.DesiredQuantity {
			allCompleted = false
			break
		}
	}

	if allCompleted {
		if err := config.DB.Model(&orderList).Update("status", "completed").Error; err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"message": "Failed to update order list status",
				"error":   err.Error(),
			})
		}
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":     "Order list updated successfully",
		"order_list":  orderList,
		"order_count": len(req.Orders),
	})
}
