package repository

import (
	"donor-service-grpc/models"

	"gorm.io/gorm"
)

type DonorRepository interface {
	Create(donor *models.Donor) error
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
