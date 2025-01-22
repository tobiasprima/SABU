package repository

import (
	"foundation-service/models"

	"gorm.io/gorm"
)

type FoundationRepository interface {
	CreateFoundation(foundation *models.Foundation) error
	GetFoundationByID(foundationID string) (*models.Foundation, error)
	AddOrderlist(orderlist *models.OrderList) error
	AddOrders(orders []models.Order) error
	GetOrderlistByID(orderlistID string, orderlist *models.OrderList) error
	GetOrdersByOrderlistID(orderlistID string, orders *[]models.Order) error
	GetOrderByID(orderID string) (*models.Order, error)
	GetOrdersArrayByOrderListID(orderListID string) ([]models.Order, error)
	UpdateOrderListStatus(orderListID string, status string) error
	GetFoundationWithEmail(foundationID string) (*models.Foundation, error)
}

type FoundationRepositoryImpl struct {
	DB *gorm.DB
}

func NewFoundationRepositoryImpl(db *gorm.DB) FoundationRepository {
	return &FoundationRepositoryImpl{db}
}

func (fr *FoundationRepositoryImpl) CreateFoundation(foundation *models.Foundation) error {
	return fr.DB.Create(foundation).Error
}

func (fr *FoundationRepositoryImpl) GetFoundationByID(foundationID string) (*models.Foundation, error) {
	foundation := new(models.Foundation)

	if err := fr.DB.Where("id = ?", foundationID).Take(foundation).Error; err != nil {
		return nil, err
	}

	return foundation, nil
}

func (fr *FoundationRepositoryImpl) AddOrderlist(orderlist *models.OrderList) error {
	return fr.DB.Create(orderlist).Error
}

func (fr *FoundationRepositoryImpl) AddOrders(orders []models.Order) error {
	return fr.DB.Create(&orders).Error
}

func (fr *FoundationRepositoryImpl) GetOrderlistByID(orderlistID string, orderlist *models.OrderList) error {
	return fr.DB.Where("id = ?", orderlistID).Take(orderlist).Error
}

func (fr *FoundationRepositoryImpl) GetOrdersByOrderlistID(orderlistID string, orders *[]models.Order) error {
	return fr.DB.Where("order_list_id = ?", orderlistID).Find(orders).Error
}

func (fr *FoundationRepositoryImpl) GetOrderByID(orderID string) (*models.Order, error) {
	var order models.Order
	if err := fr.DB.Where("id = ?", orderID).First(&order).Error; err != nil {
		return nil, err
	}
	return &order, nil
}

func (fr *FoundationRepositoryImpl) GetOrdersArrayByOrderListID(orderListID string) ([]models.Order, error) {
	var orders []models.Order
	if err := fr.DB.Where("order_list_id = ?", orderListID).Find(&orders).Error; err != nil {
		return nil, err
	}
	return orders, nil
}

func (fr *FoundationRepositoryImpl) UpdateOrderListStatus(orderListID string, status string) error {
	return fr.DB.Model(&models.OrderList{}).
		Where("id = ?", orderListID).
		Update("status", status).
		Error
}

func (fr *FoundationRepositoryImpl) GetFoundationWithEmail(foundationID string) (*models.Foundation, error) {
	var foundation models.Foundation
	if err := fr.DB.Preload("User").Where("id = ?", foundationID).First(&foundation).Error; err != nil {
		return nil, err
	}
	return &foundation, nil
}
