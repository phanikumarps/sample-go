package encode

import (
	"encoding/json"
	"log"
	"net/http"
)

func RenderJSON(w http.ResponseWriter, obj any) error {

	writeContentType(w, jsonContentType)
	w.WriteHeader(http.StatusOK)
	jb, err := json.Marshal(obj)
	if err != nil {
		return err
	}
	_, err = w.Write(jb)
	if err != nil {
		log.Println(err)
	}
	return err
}

var RenderEmptyJSON = func(w http.ResponseWriter) {
	err := RenderJSON(w, []interface{}{})
	if err != nil {
		return
	}
}

func writeContentType(w http.ResponseWriter, value string) {
	w.Header().Set("Content-Type", value)
}

var (
	jsonContentType      = "application/json; charset=utf-8"
	jsonASCIIContentType = "application/json"
)
