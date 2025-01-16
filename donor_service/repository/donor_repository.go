package repository

import (
	"donor-service/models"

	"gorm.io/gorm"
)

type DonorRepository interface {
	GetDonorByID(donorID string) (*models.Donor, error)
}

type DonorRepositoryImpl struct {
	DB *gorm.DB
}

func NewDonorRepositoryImpl(db *gorm.DB) DonorRepository {
	return &DonorRepositoryImpl{db}
}

func (dr *DonorRepositoryImpl) GetDonorByID(donorID string) (*models.Donor, error) {
	donor := new(models.Donor)

	if err := dr.DB.Where("id = ?", donorID).Take(donor).Error; err != nil {
		return nil, err
	}

	return donor, nil
}
