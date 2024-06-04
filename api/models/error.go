package models

// ResponseError ...
type ResponseError struct {
	Code    string `json:"status"`
	Message string `json:"message"`
	Data    string `json:"data"`
}
