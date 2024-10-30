package handlers

import (
	"context"
	"github.com/phanikumarps/sample-go/db"
	"github.com/phanikumarps/sample-go/encode"
	"net/http"
)

func FileHandler(ctx context.Context, store db.FileStore) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//id := r.PathValue("id")
		id := r.URL.Path[len("/file/"):]

		f, err := store.ReadFile(ctx, id)
		if err != nil {
			return
		}
		err = encode.RenderJSON(w, f)
		if err != nil {
			return
		}
	})
}
