package models

type Employee struct {
	ID         int64   `json:"id"`
	Name       string  `json:"name"`
	Occupation string  `json:"occupation"`
	Salary     float64 `json:"salary"`
}
