package handlers

import (
	"encoding/json"
	"lifeapp/internal/models"
	"lifeapp/internal/services"
	"net/http"
)

func AddTestHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var test models.Test
	err := json.NewDecoder(r.Body).Decode(&test)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = services.TestAdd(test)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
