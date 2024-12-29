package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func responseWithError(w http.ResponseWriter, code int, message string) {
	if code >= 500 {
		log.Printf("HTTP %d: %s", code, message)
	}
	responseWithJSON(w, code, map[string]string{"error": message})
}

func responseWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Fialed to marshal json response: %v", payload)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func responseWithText(w http.ResponseWriter, code int, message string) {
    w.Header().Add("Content-Type", "text/plain")
    w.WriteHeader(code)
    w.Write([]byte(message))
}