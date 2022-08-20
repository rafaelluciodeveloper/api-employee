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

func Delete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		log.Printf("Error parsing id: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	rows, err := actions.Delete(int64(id))

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
