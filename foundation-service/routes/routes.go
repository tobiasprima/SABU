package routes

import (
	"foundation-service/handlers"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo, fh *handlers.FoundationHandler) {
	e.GET("/foundation/:foundation_id", fh.GetFoundationByID)
	e.POST("/foundation/add-order/:foundation_id", fh.AddOrder)
	e.GET("/foundation/get-order/:orderlist_id", fh.GetOrder)
	e.PUT("/foundation/complete-order/:orderlist_id", fh.CompleteOrder)
}
