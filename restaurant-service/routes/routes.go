package routes

import (
	"sabu-restaurant-service/handlers"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {
	restaurantHandler := handlers.NewRestaurantHandler()

	e.POST("/restaurant/:restaurantID/add-meal", restaurantHandler.AddMeal)
}