package dtos

type DonorData struct {
	ID      string  `json:"id"`
	Name    string  `json:"name"`
	UserID  string  `json:"user_id"`
	Balance float64 `json:"balance"`
}
