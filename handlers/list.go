package handlers

import (
	"encoding/json"
	"go-api/models"
	"log"
	"net/http"
)

func List(w http.ResponseWriter, r *http.Request) {
	employes, err := models.GetAll()
	if err != nil {
		log.Printf("Erro ao obter registros %v", err)
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(employes)
}
