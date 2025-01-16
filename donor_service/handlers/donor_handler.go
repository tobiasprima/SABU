package handlers

import (
	"donor-service/dtos"
	"donor-service/repository"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type DonorHandler struct {
	DonorRepository repository.DonorRepository
}

func NewDonorHandlerImpl(donorRepository repository.DonorRepository) *DonorHandler {
	return &DonorHandler{DonorRepository: donorRepository}
}

func (dh *DonorHandler) GetDonorByID(c echo.Context) error {
	donorID := c.Param("id")

	donor, err := dh.DonorRepository.GetDonorByID(donorID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return echo.NewHTTPError(http.StatusNotFound, map[string]string{"message": "donor not found"})
		}
		return echo.NewHTTPError(http.StatusInternalServerError, map[string]string{"message": "failed to retrieve donor detail"})
	}

	res := dtos.DonorData{
		ID:      donorID,
		Name:    donor.Name,
		UserID:  donor.UserID,
		Balance: donor.Balance,
	}

	return c.JSON(http.StatusOK, res)
}
