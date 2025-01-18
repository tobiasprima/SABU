package routes

import (
	"foundation-service/handlers"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo, fh *handlers.FoundationHandler) {
	e.GET("/foundation/:id", fh.GetFoundationByID)
}
