package models

type Foundation struct {
	ID      string `json:"id" gorm:"column:id;type:uuid;default:gen_random_uuid();primaryKey"`
	Name    string `json:"name" gorm:"not null"`
	UserID  string `json:"user_id" gorm:"not null"`
	Address string `json:"address" gorm:"not null;size:100"`
}
