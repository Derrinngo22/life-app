package journal

import "github.com/go-chi/chi/v5"

func RegisterRoutes(r chi.Router, service *JournalService) {
	handler := NewJournalHandler(service)
	// r.Get("/", handler.AddJournalEntryHandler)
	r.Post("/", handler.AddJournalEntryHandler)
}
