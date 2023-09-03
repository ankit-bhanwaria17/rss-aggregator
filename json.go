package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func respondWithJSON(w http.ResponseWriter, statusCode int, payload interface{}) {

	data, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Failed to marshal JSON response: %v\n", payload)
		log.Println("Error: ", err)
		w.WriteHeader(500)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(data)
}

func respondWithErr(w http.ResponseWriter, code int, msg string) {
	if code > 499 {
		log.Println("Responding with 5xx error: ", msg)
	}

	type errResponse struct {
		Error string `json:"error"`
	}

	respondWithJSON(w, code, errResponse{
		Error: msg,
	})

}
