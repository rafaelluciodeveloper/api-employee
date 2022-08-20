package models

import uuid "github.com/satori/go.uuid"

type Employee struct {
	ID         uuid.UUID `json:"id"`
	Name       string    `json:"name"`
	Occupation string    `json:"occupation"`
	Salary     float64   `json:"salary"`
}
