package handlers

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"go-api/models/actions"
	"log"
	"net/http"
)

// Get ... Get the employee by id
// @Summary Get one employee
// @Description get employee by ID
// @Tags Employee
// @Param id path string true "Employee ID"
// @Success 200 {object} models.Employee
// @Failure 400,404 {object} object
// @Router /{id} [get]
func Get(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	employee, err := actions.Get(id)

	if err != nil {
		log.Printf("Error fetching record: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(employee)
}
