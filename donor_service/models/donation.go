package models

import "time"

type Donation struct {
	ID        string    `json:"id" gorm:"column:id;type:uuid;default:gen_random_uuid();primaryKey"`
	OrderID   string    `json:"order_id" gorm:"column:order_id;not null"`
	DonorID   string    `json:"donor_id" gorm:"column:donor_id;not null"`
	Quantity  uint      `json:"quantity" gorm:"column:quantity;not null"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at;autoCreateTime"`
}
