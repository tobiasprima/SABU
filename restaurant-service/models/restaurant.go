package models

import "time"

type Restaurant struct {
	ID        string    `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	UserID    string    `gorm:"type:uuid;not null"`
	Email     string    `gorm:"not null;unique"`
	Name      string    `gorm:"not null"`
	Address   string    `gorm:"not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}