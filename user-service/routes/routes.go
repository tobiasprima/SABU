package routes

import (
	"sabu-user-service/handlers"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes (e *echo.Echo) {
	userHandler := handlers.NewUserHandler()

	e.POST("/user/register", userHandler.RegisterUser)
	e.POST("/user/login", userHandler.LoginUser)
}