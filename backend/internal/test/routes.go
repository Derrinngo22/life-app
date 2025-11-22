package test

import "github.com/go-chi/chi/v5"

func RegisterRoutes(r chi.Router) {
	r.Post("/", AddTestHandler)
}
