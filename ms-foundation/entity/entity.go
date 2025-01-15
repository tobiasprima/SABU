package entity

type Foundation struct {
	ID      int    `json:"id" gorm:"primaryKey;autoIncrement"`
	UserID  int    `json:"user_id" gorm:"not null"`
	Address string `json:"address" gorm:"not null;size:100"`
}

type OrderList struct {
	ID           int        `json:"id" gorm:"primarykey;autoIncrement"`
	FoundationID int        `json:"foundation_id" gorm:"not null"`
	Foundation   Foundation `json:"-" gorm:"foreignkey:FoundationID"`
	Orders       []Order    `json:"orders" gorm:"foreignkey:OrderListID"`
	Status       string     `json:"status"`
}

type Order struct {
	ID              int `json:"id" gorm:"primarykey;autoIncrement"`
	OrderListID     int `json:"order_list_id" gorm:"not null"`
	MealsID         int `json:"meals_id" gorm:"not null"`
	Quantity        int `json:"quantity" gorm:"not null;default:0"`
	DesiredQuantity int `json:"desired_quantity" gorm:"not null"`
}

type OrderRequest struct {
	OrderListID int `json:"orderlist_id" validate:"required"`
	Orders      []struct {
		MealsID         int `json:"meals_id" validate:"required"`
		DesiredQuantity int `json:"quantity" validate:"required"`
	} `json:"orders" validate:"required"`
}
