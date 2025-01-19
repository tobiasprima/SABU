package routes

import (
	"sabu-restaurant-service/handlers"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {
	restaurantHandler := handlers.NewRestaurantHandler()

	e.POST("/restaurant/add-meal/:restaurant_id", restaurantHandler.AddMeal)
	e.GET("/restaurant/get-meals/:restaurant_id", restaurantHandler.GetMealsByRestaurantID)
	e.GET("/restaurant/get-meal/:meal_id", restaurantHandler.GetMealByID)
	e.PATCH("/restaurant/update-meal/:meal_id", restaurantHandler.UpdateMeal)
	e.DELETE("/restaurant/delete-meal/:meal_id", restaurantHandler.DeleteMeal)
}