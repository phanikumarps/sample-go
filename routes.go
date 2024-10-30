package main

import (
	"context"
	"github.com/phanikumarps/sample-go/db"
	"github.com/phanikumarps/sample-go/handlers"
	"net/http"
)

func addRoutes(ctx context.Context, mux *http.ServeMux, store *db.Store) {

	mux.Handle("GET /__health", handlers.HealthHandler())
	mux.Handle("GET /file/{id}", handlers.FileHandler(ctx, store))
}
