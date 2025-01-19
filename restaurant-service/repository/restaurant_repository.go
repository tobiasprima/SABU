package repository

import (
	"context"
	"errors"
	"sabu-restaurant-service/config"
	"sabu-restaurant-service/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func (r *RestaurantRepository) GetRestaurantByID(restaurantID string) (*models.Restaurant, error) {
	var restaurant models.Restaurant

	err := r.DB.Where("ID = ?", restaurantID).First(&restaurant).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &restaurant, nil
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

func (r *RestaurantRepository) UpdateMeal(ctx context.Context, mealId string, updates bson.M) error {
	objectID, err := primitive.ObjectIDFromHex(mealId)
	if err != nil {
		return errors.New("Invalid meal ID format")
	}

	filter := bson.M{"_id": objectID}
	update := bson.M{"$set": updates}

	_, err = r.MealsCollection.UpdateOne(ctx, filter, update)
	return err
}

func (r *RestaurantRepository) DeleteMeal(ctx context.Context, mealId string) error {
	objectID, err := primitive.ObjectIDFromHex(mealId)
	if err != nil {
		return errors.New("Invalid meal ID format")
	}

	filter := bson.M{"_id": objectID}
	result, err := r.MealsCollection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}

	if result.DeletedCount == 0 {
		return mongo.ErrNoDocuments
	}

	return nil
}