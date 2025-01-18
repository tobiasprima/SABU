package models

type Order struct {
	ID              int `json:"id" gorm:"primarykey;autoIncrement"`
	OrderListID     int `json:"order_list_id" gorm:"not null"`
	MealsID         int `json:"meals_id" gorm:"not null"`
	Quantity        int `json:"quantity" gorm:"not null;default:0"`
	DesiredQuantity int `json:"desired_quantity" gorm:"not null"`
}
