package models

type Foundation struct {
	ID      string `json:"id" gorm:"column:id;type:uuid;default:gen_random_uuid();primaryKey"`
	Name    string `json:"name" gorm:"not null"`
	UserID  string `json:"user_id" gorm:"not null"`
	User    User   `json:"-" gorm:"foreignkey:UserID"`
	Address string `json:"address" gorm:"not null;size:100"`
}

type User struct {
	ID       string `json:"id" gorm:"column:id;type:uuid;default:gen_random_uuid();primaryKey"`
	Email    string `json:"email" gorm:"not null;unique"`
	Password string `json:"password" gorm:"not null"`
	UserType string `json:"user_type" gorm:"not null"`
}
