package dtos

type OrderRequest struct {
	Orders       []struct {
		MealsID         string `json:"meals_id"`
		DesiredQuantity int    `json:"quantity"`
	} `json:"orders"`
}
