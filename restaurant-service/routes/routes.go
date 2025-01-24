package routes

import (
	_ "sabu-restaurant-service/docs"
	"sabu-restaurant-service/handlers"

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func RegisterRoutes(e *echo.Echo) {
	restaurantHandler := handlers.NewRestaurantHandler()

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.GET("/restaurants", restaurantHandler.GetRestaurants)
	e.GET("/restaurant/:restaurant_id", restaurantHandler.GetRestaurantByID)
	e.GET("/restaurant/get-meals/:restaurant_id", restaurantHandler.GetMealsByRestaurantID)
	e.GET("/restaurant/get-meal/:meal_id", restaurantHandler.GetMealByID)
	e.POST("/restaurant/add-meal/:restaurant_id", restaurantHandler.AddMeal)
	e.PATCH("/restaurant/update-meal/:meal_id", restaurantHandler.UpdateMeal)
	e.DELETE("/restaurant/delete-meal/:meal_id", restaurantHandler.DeleteMeal)
}
