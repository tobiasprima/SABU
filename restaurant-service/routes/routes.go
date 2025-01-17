package routes

import (
	"sabu-restaurant-service/handlers"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {
	restaurantHandler := handlers.NewRestaurantHandler()

	e.POST("/restaurant/:restaurant_id/add-meal", restaurantHandler.AddMeal)
	e.GET("/restaurant/:restaurant_id/get-meals", restaurantHandler.GetMealsByRestaurantID)
}