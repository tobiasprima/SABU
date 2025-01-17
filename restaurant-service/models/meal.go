package models

type Meal struct {
	ID           string  `json:"id" bson:"_id,omitempty"`
	RestaurantID string  `json:"restaurant_id" bson:"restaurant_id"`
	Name         string  `json:"name" bson:"name"`
	Description  string  `json:"description" bson:"description"`
	Price        float64 `json:"price" bson:"price"`
	Stock        int     `json:"stock" bson:"stock"`
}