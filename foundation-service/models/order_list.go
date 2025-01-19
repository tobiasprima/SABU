package models

type OrderList struct {
	ID           string     `json:"id" gorm:"primarykey;autoIncrement"`
	FoundationID string     `json:"foundation_id" gorm:"not null"`
	Foundation   Foundation `json:"-" gorm:"foreignkey:FoundationID"`
	Status       string     `json:"status"`
}
