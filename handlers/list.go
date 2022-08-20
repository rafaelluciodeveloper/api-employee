package handlers

import (
	"encoding/json"
	"go-api/models/actions"
	"log"
	"net/http"
)

func List(w http.ResponseWriter, r *http.Request) {
	employes, err := actions.GetAll()
	if err != nil {
		log.Printf("Error getting records %v", err)
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(employes)
}
