package handlers

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"go-api/models"
	"go-api/models/actions"
	"log"
	"net/http"
)

// Delete ... Delete the employee by id
// @Summary Delete one employee
// @Description delete employee by ID
// @Tags Employee
// @Param id path string true "Employee ID"
// @Success 200 {object} models.Employee
// @Failure 400,404 {object} object
// @Router /{id} [delete]
func Delete(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	rows, err := actions.Delete(id)

	if err != nil {
		log.Printf("Error deleting record: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	if rows > 1 {
		log.Printf("Error: %d records were removed", rows)
	}

	resp := models.Response{
		Error:   false,
		Message: "Successfully deleted record!",
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
