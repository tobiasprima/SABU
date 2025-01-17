package repository

import (
	"context"
	"sabu-restaurant-service/config"
	"sabu-restaurant-service/models"

	"go.mongodb.org/mongo-driver/bson"
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

func (r *RestaurantRepository) GetMealsByRestaurantID(ctx context.Context, restaurantID string) ([]models.Meal, error) {
	cursor, err := r.MealsCollection.Find(ctx, bson.M{"restaurant_id": restaurantID})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var meals []models.Meal
	if err := cursor.All(ctx, &meals); err != nil {
		return nil, err
	}

	return meals, nil
}