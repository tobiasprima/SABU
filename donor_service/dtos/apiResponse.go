package dtos

type APISuccessResponse struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type APIErrorResponse struct {
	Status int         `json:"status"`
	Type   string      `json:"type"`
	Detail interface{} `json:"detail"`
}
