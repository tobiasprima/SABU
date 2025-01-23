package repository

import (
	"context"
	"errors"
	"restaurant-service-grpc/config"
	"restaurant-service-grpc/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

type RestaurantRepository struct {
	DB              *gorm.DB
	MealsCollection *mongo.Collection
}

func NewRestaurantRepository() *RestaurantRepository {
	return &RestaurantRepository{
		DB:              config.Database.Restaurant,
		MealsCollection: config.MealsCollection,
	}
}

func (r *RestaurantRepository) CreateRestaurant(restaurant *models.Restaurant) error {
	return r.DB.Create(restaurant).Error
}

func (r *RestaurantRepository) GetMealByID(ctx context.Context, mealId string) (*models.Meal, error) {
	objectID, err := primitive.ObjectIDFromHex(mealId)
	if err != nil {
		return nil, errors.New("Invalid meal ID format")
	}

	var meal models.Meal
	err = r.MealsCollection.FindOne(ctx, bson.M{"_id": objectID}).Decode(&meal)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}

	return &meal, nil
}

func (r *RestaurantRepository) DeductMealStock(ctx context.Context, meal *models.Meal) error {
	objectID, err := primitive.ObjectIDFromHex(meal.ID)
	if err != nil {
		return errors.New("invalid meal ID format")
	}

	filter := bson.M{"_id": objectID}
	update := bson.M{"$inc": bson.M{"stock": -meal.Stock}}

	result, err := r.MealsCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	if result.ModifiedCount == 0 {
		return errors.New("not found")
	}

	return nil
}