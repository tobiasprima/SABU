package models

type OrderList struct {
	ID           int        `json:"id" gorm:"primarykey;autoIncrement"`
	FoundationID int        `json:"foundation_id" gorm:"not null"`
	Foundation   Foundation `json:"-" gorm:"foreignkey:FoundationID"`
	Orders       []Order    `json:"orders" gorm:"foreignkey:OrderListID"`
	Status       string     `json:"status"`
}
