package main

import (
	"log"
	"net/http"

	"lifeapp/internal/db"
	"lifeapp/internal/handlers"
)

func main() {
	db.Connect()

	mux := http.NewServeMux()
	mux.HandleFunc("/tests", handlers.AddTestHandler)
	// mux.HandleFunc("/tests/", handlers.TestGet)

	log.Println("Server running on :8080")

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal(err)
	}
}
