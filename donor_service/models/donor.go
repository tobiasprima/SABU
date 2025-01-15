package models

type Donor struct {
	ID          uint          `json:"id" gorm:"primaryKey;column:id;autoIncrement"`
	Name        string        `json:"name" gorm:"column:name;not null"`
	UserID      uint          `json:"user_id" gorm:"column:user_id;not null"`
	Balance     float64       `json:"balance" gorm:"column:balance;default:0"`
	TopUp       []TopUp       `gorm:"foreignKey:donor_id;references:id"`
	Transaction []Transaction `gorm:"foreignKey:donor_id;references:id"`
}
