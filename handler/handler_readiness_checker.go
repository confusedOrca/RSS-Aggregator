package handler

import (
	"net/http"
)

func HandlerReadinessChecker(responseWriter http.ResponseWriter, request *http.Request) {
	respondWithJSON(responseWriter, 200, struct{}{})
}
