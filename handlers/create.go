package handlers

import (
	"encoding/json"
	"fmt"
	"go-api/models"
	"go-api/models/actions"
	"log"
	"net/http"
)

func Create(w http.ResponseWriter, r *http.Request) {
	var employe models.Employee
	err := json.NewDecoder(r.Body).Decode(&employe)
	if err != nil {
		log.Printf("Erro ao fazer decode do json: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	id, err := actions.Insert(employe)
	var resp models.Response

	if err != nil {
		resp = models.Response{
			Error:   true,
			Message: fmt.Sprintf("An error occurred while trying to enter %v", err),
		}
	} else {
		resp = models.Response{
			Error:   false,
			Message: fmt.Sprintf("Record inserted successfully ID: %v", id),
		}
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
