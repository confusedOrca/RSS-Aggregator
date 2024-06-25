package main

import (
	"net/http"
)

func handlerReadinessChecker(responseWriter http.ResponseWriter, request *http.Request) {
	respondWithJSON(responseWriter, 200, struct{}{})
}
