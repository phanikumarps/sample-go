package main

import (
	"context"
	"github.com/phanikumarps/sample-go/db"
	"net/http"
)

func run(ctx context.Context) error {

	store, err := db.Connect(ctx, &db.Config{
		Host: "test.pgx.com",
		Port: 5432,
		User: "postgres",
		Pass: "test123",
		Name: "test",
	})
	if err != nil {
		return err
	}
	mux := http.NewServeMux()
	addRoutes(ctx, mux, store)
	err = http.ListenAndServe("0.0.0.0:5001", mux)
	if err != nil {
		return err
	}
	return nil
}
