package routes

import (
	_ "sabu-user-service/docs"
	"sabu-user-service/handlers"
	"sabu-user-service/proto/pb"

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func RegisterRoutes(e *echo.Echo, restaurantClient pb.RestaurantServiceClient, donorClient pb.DonorServiceClient, foundationClient pb.FoundationServiceClient) {
	userHandler := handlers.NewUserHandler(restaurantClient, donorClient, foundationClient)

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.POST("/user/register", userHandler.RegisterUser)
	e.POST("/user/login", userHandler.LoginUser)
}
