package repository

import (
	"sabu-restaurant-service/config"
	"sabu-restaurant-service/models"

	"gorm.io/gorm"
)

type RestaurantRepository struct {
	DB *gorm.DB
}

func NewRestaurantRepository() *RestaurantRepository {
	return &RestaurantRepository{DB: config.Database.Restaurant}
}

func (r *RestaurantRepository) CreateRestaurant(restaurant *models.Restaurant) error {
	return r.DB.Create(restaurant).Error
}