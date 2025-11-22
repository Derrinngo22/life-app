package journal

import (
	"context"
	"encoding/json"
	"net/http"
)

type JournalHandler struct {
	service *JournalService
}

func NewJournalHandler(service *JournalService) *JournalHandler {
	return &JournalHandler{service: service}
}

func (h *JournalHandler) AddJournalEntryHandler(w http.ResponseWriter, r *http.Request) {

	var entry JournalEntry

	err := json.NewDecoder(r.Body).Decode(&entry)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.service.CreateEntry(context.Background(), &entry)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(entry)
}
