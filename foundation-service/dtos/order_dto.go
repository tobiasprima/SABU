package dtos

type OrderRequest struct {
	FoundationID int `json:"yayasan_id"`
	Orders       []struct {
		MealsID         int `json:"meals_id"`
		DesiredQuantity int `json:"quantity"`
	} `json:"orders"`
}
