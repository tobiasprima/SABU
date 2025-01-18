package dtos

type AddMealRequest struct {
	Name			string		`json:"name"`
	Description		string		`json:"description"`
	Price			float64		`json:"price"`
	Stock			int			`json:"stock"`
}