package handlers

import (
	"foundation-service/dtos"
	"foundation-service/repository"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type FoundationHandler struct {
	FoundationRepository repository.FoundationRepository
}

func NewFoundationHandlerImpl(foundationRepository repository.FoundationRepository) *FoundationHandler {
	return &FoundationHandler{FoundationRepository: foundationRepository}
}

func (fh *FoundationHandler) GetFoundationByID(c echo.Context) error {
	foundationID := c.Param("id")

	foundation, err := fh.FoundationRepository.GetFoundationByID(foundationID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return echo.NewHTTPError(http.StatusNotFound, map[string]string{"message": "foundation not found"})
		}
		return echo.NewHTTPError(http.StatusInternalServerError, map[string]string{"message": "failed to retrieve foundation detail"})
	}

	res := dtos.FoundationData{
		ID:      foundationID,
		Name:    foundation.Name,
		UserID:  foundation.UserID,
		Address: foundation.Address,
	}

	return c.JSON(http.StatusOK, res)
}
