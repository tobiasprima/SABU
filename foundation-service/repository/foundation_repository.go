package repository

import (
	"foundation-service/models"

	"gorm.io/gorm"
)

type FoundationRepository interface {
	Create(foundation *models.Foundation) error
	GetFoundationByID(foundationID string) (*models.Foundation, error)
}

type FoundationRepositoryImpl struct {
	DB *gorm.DB
}

func NewFoundationRepositoryImpl(db *gorm.DB) FoundationRepository {
	return &FoundationRepositoryImpl{db}
}

func (fr *FoundationRepositoryImpl) Create(foundation *models.Foundation) error {
	return fr.DB.Create(foundation).Error
}

func (fr *FoundationRepositoryImpl) GetFoundationByID(foundationID string) (*models.Foundation, error) {
	foundation := new(models.Foundation)

	if err := fr.DB.Where("id = ?", foundationID).Take(foundation).Error; err != nil {
		return nil, err
	}

	return foundation, nil
}
