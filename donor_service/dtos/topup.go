package dtos

import "time"

type TopUpRequest struct {
	Amount float64 `json:"amount"`
}

type XenditWebhookRequest struct {
	ExternalID    string    `json:"external_id"`
	InvoiceID     string    `json:"invoice_id"`
	PaymentMethod string    `json:"payment_method"`
	Status        string    `json:"status"`
	CompletedAt   time.Time `json:"paid_at"`
}
