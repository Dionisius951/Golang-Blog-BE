package models

type ApiResponse struct {
	Success bool `json:"success"`
	Message string `json:"message"`
	Error error `json:"error,omitempty"`
	Data interface{} `json:"data,omitempty"`
}