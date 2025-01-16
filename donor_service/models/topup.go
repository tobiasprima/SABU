package models

import "time"

type TopUp struct {
	ID        string    `json:"id" gorm:"column:id;type:uuid;default:gen_random_uuid();primaryKey"`
	DonorID   string    `json:"donor_id" gorm:"column:donor_id;not null"`
	Amount    float64   `json:"amount" gorm:"column:amount;not null"`
	Status    string    `json:"status" gorm:"column:status;size:50;not null"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at;autoCreateTime"`
}
