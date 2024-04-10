package server

import "github.com/go-chi/chi/v5"

func newServer() *chi.Mux {
	r := chi.NewRouter()
	return r
}
