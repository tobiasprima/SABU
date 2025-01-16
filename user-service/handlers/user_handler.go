package handlers

import (
	"net/http"
	"sabu-user-service/models"
	"sabu-user-service/repository"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	UserRepo       *repository.UserRepository
}

func NewUserHandler() *UserHandler {
	return &UserHandler{
		UserRepo:       repository.NewUserRepository(),
	}
}

func (h *UserHandler) RegisterUser(c echo.Context) error {
	user := new(models.User)
	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
	}

	// Save user to the database
	if err := h.UserRepo.CreateUser(user); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to register user"})
	}

	// Handle restaurant-specific logic
	// if user.UserType == "restaurant" {
	// 	restaurant := &models.Restaurant{
	// 		UserID:  user.ID,
	// 		Name:    c.FormValue("restaurant_name"),
	// 		Address: c.FormValue("address"),
	// 	}
	// 	if err := h.RestaurantRepo.CreateRestaurant(restaurant); err != nil {
	// 		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create restaurant"})
	// 	}
	// }

	return c.JSON(http.StatusCreated, map[string]string{"message": "User registered successfully"})
}
