package routes

import (
	"donor-service/handlers"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo, dh handlers.DonorHandler) {
	e.GET("/donor/:id", dh.GetDonorByID)
	e.POST("/donor/top-up/:donorID", dh.TopUp)

	// Webhook endpoint
	e.POST("/donor/update-top-up", dh.UpdateTopUpStatus)
}
