package routes

import (
	"sabu-user-service/handlers"
	"sabu-user-service/proto/pb"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes (e *echo.Echo, restaurantClient pb.RestaurantServiceClient) {
	userHandler := handlers.NewUserHandler(restaurantClient)

	e.POST("/user/register", userHandler.RegisterUser)
	e.POST("/user/login", userHandler.LoginUser)
}