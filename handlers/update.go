package handlers

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"go-api/models"
	"go-api/models/actions"
	"log"
	"net/http"
)

// Update ... Update employee
// @Summary Update employee based on parameters
// @Description Update employee
// @Tags Employee
// @Accept json
// @Param user body models.Employee true "Employee Data"
// @Success 200 {object} object
// @Failure 400,500 {object} object
// @Router / [put]
func Update(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	var employe models.Employee
	err := json.NewDecoder(r.Body).Decode(&employe)
	if err != nil {
		log.Printf("Error decoding json: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	rows, err := actions.Update(id, employe)

	if err != nil {
		log.Printf("Error updating record: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	if rows > 1 {
		log.Printf("Error: %d records were updated", rows)
	}

	resp := models.Response{
		Error:   false,
		Message: "Record updated successfully!",
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
