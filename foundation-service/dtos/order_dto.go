package dtos

type OrderRequest struct {
	OrderlistID  string `json:"orderlist_id"`
	FoundationID string `json:"foundation_id"`
	Orders       []struct {
		MealsID         string `json:"meals_id"`
		DesiredQuantity int    `json:"quantity"`
	} `json:"orders"`
}
