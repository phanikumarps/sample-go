package handlers

import (
	"github.com/phanikumarps/sample-go/encode"
	"net/http"
	"time"
)

func HealthHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		type health struct {
			Now    time.Time `json:"time,omitempty"`
			Status string    `json:"status,omitempty"`
		}
		var hl health
		hl.Now = time.Now().Local().UTC()
		hl.Status = "OK"

		err := encode.RenderJSON(w, hl)
		if err != nil {
			return
		}

	})
}
