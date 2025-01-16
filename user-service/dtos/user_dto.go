package dtos

type RegisterUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	UserType string `json:"user_type"`
	Name     string `json:"name,omitempty"`    // Optional for non-restaurant users
	Address  string `json:"address,omitempty"` // Optional for non-restaurant users

}