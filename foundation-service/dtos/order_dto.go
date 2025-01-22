package dtos

type OrderlistRequest struct {
	FoundationID string `json:"foundation_id"`
}

type OrderRequest struct {
	OrderlistID  string `json:"orderlist_id"`
	Orders       []struct {
		MealsID         string `json:"meals_id"`
		DesiredQuantity int    `json:"quantity"`
	} `json:"orders"`
}
