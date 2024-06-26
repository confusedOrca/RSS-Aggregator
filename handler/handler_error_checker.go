package handler

import (
	"net/http"
)

func HandlerErrorChecker(responseWriter http.ResponseWriter, request *http.Request) {
	respondWithError(responseWriter, 400, "Something Went Wrong!")
}
