package models

import "time"

type TopUp struct {
	ID        uint      `json:"id" gorm:"primaryKey;column:id;autoIncrement"`
	DonorID   uint      `json:"donor_id" gorm:"column:donor_id;not null"`
	Amount    float64   `json:"amount" gorm:"column:amount;not null"`
	Status    string    `json:"status" gorm:"column:status;not null"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at;autoCreateTime"`
}
