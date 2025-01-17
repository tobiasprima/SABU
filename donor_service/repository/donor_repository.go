package repository

import (
	"donor-service/models"

	"gorm.io/gorm"
)

type DonorRepository interface {
	Create(donor *models.Donor) error
	GetDonorByID(donorID string) (*models.Donor, error)
}

type DonorRepositoryImpl struct {
	DB *gorm.DB
}

func NewDonorRepositoryImpl(db *gorm.DB) DonorRepository {
	return &DonorRepositoryImpl{db}
}

func (dr *DonorRepositoryImpl) Create(donor *models.Donor) error {
	return dr.DB.Create(donor).Error
}

func (dr *DonorRepositoryImpl) GetDonorByID(donorID string) (*models.Donor, error) {
	donor := new(models.Donor)

	if err := dr.DB.Where("id = ?", donorID).Take(donor).Error; err != nil {
		return nil, err
	}

	return donor, nil
}
