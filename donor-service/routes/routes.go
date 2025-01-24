package routes

import (
	_ "donor-service/docs"
	"donor-service/handlers"

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func RegisterRoutes(e *echo.Echo, dh handlers.DonorHandler) {
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.GET("/donor/:id", dh.GetDonorByID)
	e.POST("/donor/top-up/:donorID", dh.TopUp)
	e.GET("/donor/top-up-history/:donorID", dh.GetTopUpHistory)
	e.POST("/donor/donate/:donorID", dh.Donate)
	e.GET("/donor/donation-history/:donorID", dh.GetDonationHistory)

	// Webhook endpoint
	e.POST("/donor/update-top-up", dh.UpdateTopUpStatus)
}
