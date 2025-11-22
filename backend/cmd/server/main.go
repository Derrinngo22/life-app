package main

import (
	"log"
	"net/http"

	"lifeapp/internal/db"
	"lifeapp/internal/journal"
	"lifeapp/internal/test"

	"github.com/go-chi/chi/v5"
)

func main() {
	db.Connect()

	r := chi.NewRouter()

	journalService := journal.NewJournalService(journal.NewJournalRepository())

	features := map[string]func(chi.Router){
		"/journal": func(r chi.Router) { journal.RegisterRoutes(r, journalService) },
		"/test":    func(r chi.Router) { test.RegisterRoutes(r) },
	}

	for prefix, register := range features {
		r.Route(prefix, register)
	}

	log.Println("Server running on :8080")

	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatal(err)
	}
}
