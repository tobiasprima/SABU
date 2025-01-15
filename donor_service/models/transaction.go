package models

import "time"

type Transaction struct {
	ID        uint      `json:"id" gorm:"primaryKey;column:id;autoIncrement"`
	OrderID   uint      `json:"order_id" gorm:"column:order_id;not null"`
	DonorID   uint      `json:"donor_id" gorm:"column:donor_id;not null"`
	Quantity  uint      `json:"quantity" gorm:"column:quantity;not null"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at;autoCreateTime"`
}
