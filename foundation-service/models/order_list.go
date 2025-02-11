package models

type OrderList struct {
	ID           string     `json:"id" gorm:"column:id;type:uuid;default:gen_random_uuid();primaryKey"`
	FoundationID string     `json:"foundation_id" gorm:"not null"`
	Foundation   Foundation `json:"-" gorm:"foreignkey:FoundationID"`
	Status       string     `json:"status"`
}
