package models

type Foundation struct {
	ID      string `json:"id" gorm:"primaryKey;autoIncrement"`
	Name    string `json:"name" gorm:"not null"`
	UserID  string `json:"user_id" gorm:"not null"`
	Address string `json:"address" gorm:"not null;size:100"`
}
