package handlers

import (
	"encoding/json"
	"fmt"
	"go-api/models"
	"go-api/models/actions"
	"log"
	"net/http"
)

// Create ... Create employee
// @Summary Create new employee based on paramters
// @Description Create new employee
// @Tags Employee
// @Accept json
// @Param user body models.Employee true "Employee Data"
// @Success 200 {object} object
// @Failure 400,500 {object} object
// @Router / [post]
func Create(w http.ResponseWriter, r *http.Request) {
	var employee models.Employee
	err := json.NewDecoder(r.Body).Decode(&employee)
	if err != nil {
		log.Printf("Erro ao fazer decode do json: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	id, err := actions.Insert(employee)
	var resp models.Response

	if err != nil {
		resp = models.Response{
			Error:   true,
			Message: fmt.Sprintf("An error occurred while trying to enter %v", err),
		}
	} else {
		resp = models.Response{
			Error:   false,
			Message: fmt.Sprintf("Record inserted successfully ID: %s", id),
		}
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
