package models

type Order struct {
	ID              string `json:"id" gorm:"primarykey;autoIncrement"`
	OrderListID     string `json:"order_list_id" gorm:"not null"`
	MealsID         string `json:"meals_id" gorm:"not null"`
	Quantity        int    `json:"quantity" gorm:"not null;default:0"`
	DesiredQuantity int    `json:"desired_quantity" gorm:"not null"`
}
