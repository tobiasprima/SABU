package repository

import (
	"context"
	"sabu-restaurant-service/config"
	"sabu-restaurant-service/models"

	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

type RestaurantRepository struct {
	DB *gorm.DB
	MealsCollection *mongo.Collection
}

func NewRestaurantRepository() *RestaurantRepository {
	return &RestaurantRepository{
		DB: config.Database.Restaurant,
		MealsCollection: config.MealsCollection,
	}
}

func (r *RestaurantRepository) CreateRestaurant(restaurant *models.Restaurant) error {
	return r.DB.Create(restaurant).Error
}

func (r *RestaurantRepository) AddMeal(ctx context.Context, meals *models.Meal) error {
	_, err := r.MealsCollection.InsertOne(ctx, meals)
	return err
}