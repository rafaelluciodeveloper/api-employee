package models

type Response struct {
	Error   bool   `json:"Error"`
	Message string `json:"Message"`
}
