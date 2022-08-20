package handlers

import (
	"encoding/json"
	"go-api/models/actions"
	"log"
	"net/http"
)

// List ... Get all employees
// @Summary Get all employees
// @Description get all employees
// @Tags Employee
// @Success 200 {array} models.Employee
// @Failure 404 {object} object
// @Router / [get]
func List(w http.ResponseWriter, r *http.Request) {
	employees, err := actions.GetAll()
	if err != nil {
		log.Printf("Error getting records %v", err)
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(employees)
}
