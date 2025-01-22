package dtos

type DonateRequest struct {
	OrderID  string `json:"order_id"`
	Quantity uint32 `json:"quantity"`
}
