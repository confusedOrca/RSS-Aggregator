package main

import (
	"net/http"
)

func handlerErrorChecker(responseWriter http.ResponseWriter, request *http.Request) {
	respondWithError(responseWriter, 400, "Something Went Wrong!")
}
