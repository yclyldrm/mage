package utils

import (
	"encoding/json"
	"log"
	"net/http"
)

func RenderJSON(w http.ResponseWriter, statusCode int, params interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(statusCode)
	if err := json.NewEncoder(w).Encode(params); err != nil {
		log.Printf("ERROR: renderJson - %q\n", err)
	}
}
