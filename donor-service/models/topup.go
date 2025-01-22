package models

import "time"

type TopUp struct {
	ID            string     `json:"id" gorm:"column:id;type:uuid;default:gen_random_uuid();primaryKey"`
	DonorID       string     `json:"donor_id" gorm:"column:donor_id;not null"`
	Amount        float64    `json:"amount" gorm:"column:amount;not null"`
	InvoiceID     string     `json:"invoice_id" gorm:"column:invoice_id;not null"`
	InvoiceUrl    string     `json:"invoice_url" gorm:"column:invoice_url;not null"`
	PaymentMethod string     `json:"payment_method" gorm:"column:payment_method;size:50;not null"`
	Status        string     `json:"status" gorm:"column:status;size:50;not null"`
	CreatedAt     time.Time  `json:"created_at" gorm:"column:created_at;autoCreateTime"`
	CompletedAt   *time.Time `json:"completed_at,omitempty" gorm:"column:completed_at"`
}

type Invoice struct {
	InvoiceID  string  `json:"invoice_id"`
	InvoiceUrl string  `json:"invoice_url"`
	Amount     float64 `json:"amount"`
	Method     string  `json:"method"`
	Status     string  `json:"status"`
}
