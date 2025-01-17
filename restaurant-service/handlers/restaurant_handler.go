package handlers

import (
	"net/http"
	"sabu-restaurant-service/dtos"
	"sabu-restaurant-service/models"
	"sabu-restaurant-service/repository"

	"github.com/labstack/echo/v4"
)

type RestaurantHandler struct {
	RestaurantRepo		*repository.RestaurantRepository
}

func NewRestaurantHandler() *RestaurantHandler {
	return &RestaurantHandler{
		RestaurantRepo: repository.NewRestaurantRepository(),
	}
}

func (h *RestaurantHandler) AddMeal(c echo.Context) error {
	ctx := c.Request().Context()
	req := new(dtos.AddMealRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
	}

	restaurantID := c.Param("restaurantID")
	if restaurantID == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Restaurant ID is required"})
	}

	if req.Name == "" || req.Price <= 0 || req.Stock < 0 {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Missing or invalid fields (name, price, stock)",
		})
	}

	meal := &models.Meal{
		RestaurantID: restaurantID,
		Name: req.Name,
		Description: req.Description,
		Price: req.Price,
		Stock: req.Stock,
	}
	err := h.RestaurantRepo.AddMeal(ctx, meal)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to add meal"})
	}
	return c.JSON(http.StatusCreated, map[string]string{"message": "Meal added successfully"})
}