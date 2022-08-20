package handlers

import (
	"encoding/json"
	"fmt"
	"go-api/models"
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
	id, err := models.Insert(employe)
	var resp map[string]any

	if err != nil {
		resp = map[string]any{
			"Error":   true,
			"Message": fmt.Sprintf("Ocorreu um erro ao tentar inserir %v", err),
		}
	} else {
		resp = map[string]any{
			"Error":   false,
			"Message": fmt.Sprintf("Funcionario inserido com sucesso! ID: %v", id),
		}
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
