package main

import (
	"github.com/aadi-1024/gotasks-web/pkg/handlers"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func NewRouter(handlerRepo *handlers.Repository) http.Handler {
	mux := chi.NewMux()

	mux.Get("/tasks", handlerRepo.GetAllTasks)
	mux.Get("/tasks/{id}", handlerRepo.GetById)

	return mux
}
