package handler

import (
	"encoding/json"
	"log"
	"net/http"
)

type errResponse struct {
	Error string `json:"error"`
}

func isServerError(code int) bool {
	return code >= 500 && code <= 599
}

func respondWithJSON(responseWriter http.ResponseWriter, responseCode int, payload interface{}) {
	data, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Failed to marshal JSON response: %v", payload)
		responseWriter.WriteHeader(500)
		return
	}

	responseWriter.Header().Add("Content-Type", "application/json")
	responseWriter.WriteHeader(responseCode)
	responseWriter.Write(data)
}

func respondWithError(responseWriter http.ResponseWriter, errorCode int, errorMsg string) {
	if isServerError(errorCode) {
		log.Println("Responding with a 5XX error: ", errorMsg)
	}

	respondWithJSON(responseWriter, errorCode, errResponse{
		Error: errorMsg,
	})
}
