package models

import "time"

type User struct {
	ID        string    `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"` // Automatically generated UUID
	Email     string    `gorm:"unique;not null" json:"email"`                            // Unique and required
	Password  string    `gorm:"not null" json:"-"`                                       // Required, omit in JSON responses
	UserType  string    `gorm:"not null" json:"user_type"`                               // Required, e.g., restaurant, donor, foundation
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`                        // Automatically set on creation
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`                        // Automatically set on update
}
