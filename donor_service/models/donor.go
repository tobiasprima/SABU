package models

import "time"

type Donor struct {
	ID          string        `json:"id" gorm:"column:id;type:uuid;default:gen_random_uuid();primaryKey"`
	Name        string        `json:"name" gorm:"column:name;size:50;not null"`
	UserID      string        `json:"user_id" gorm:"column:user_id;type:uuid;not null"`
	Balance     float64       `json:"balance" gorm:"column:balance;default:0"`
	CreatedAt   time.Time     `json:"created_at" gorm:"column:created_at;autoCreateTime"`
	UpdatedAt   time.Time     `json:"updated_at" gorm:"column:updated_at;autoUpdateTime"`
	TopUp       []TopUp       `gorm:"foreignKey:donor_id;references:id"`
	Transaction []Transaction `gorm:"foreignKey:donor_id;references:id"`
}
