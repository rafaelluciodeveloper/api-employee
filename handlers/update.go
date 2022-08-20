package handlers

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"go-api/models"
	"go-api/models/actions"
	"log"
	"net/http"
	"strconv"
)

func Update(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		log.Printf("Erro ao fazer parse do id: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	var employe models.Employee
	err = json.NewDecoder(r.Body).Decode(&employe)
	if err != nil {
		log.Printf("Error decoding json: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	rows, err := actions.Update(int64(id), employe)

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
