package repository

import (
	"errors"
	"foundation-service/models"

	"gorm.io/gorm"
)

type FoundationRepository interface {
	Create(foundation *models.Foundation) error
	GetFoundationByID(foundationID string) (*models.Foundation, error)
	GetOrderByID(orderID string) (*models.Order, error)
	AddOrderQuantity(orderID string, quantity int) error
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

func (fr *FoundationRepositoryImpl) GetOrderByID(orderID string) (*models.Order, error) {
	order := new(models.Order)

	if err := fr.DB.Where("id = ?", orderID).Take(order).Error; err != nil {
		return nil, err
	}

	return order, nil
}

func (fr *FoundationRepositoryImpl) AddOrderQuantity(orderID string, quantity int) error {
	order := models.Order{ID: orderID}

	res := fr.DB.Model(&order).Update("quantity", gorm.Expr("quantity + ?", quantity))
	if err := res.Error; err != nil {
		return err
	}

	if res.RowsAffected == 0 {
		return errors.New("not found")
	}

	return nil
}
