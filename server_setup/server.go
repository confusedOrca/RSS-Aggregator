package server_setup

import (
	"log"
	"net/http"
)

func StartServer(handler http.Handler, port string) {
	server := &http.Server{
		Handler: handler,
		Addr:    ":" + port,
	}

	log.Printf("Server starting on port %v", port)
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
