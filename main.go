package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/confusedOrca/RSS-Aggregator/internal/database"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"

	_ "github.com/lib/pq"
)

type apiConfig struct {
	DB *database.Queries
}

func main() {
	godotenv.Load(".env")
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("Port not found in environment")
	}

	dbUrl := os.Getenv("DB_URL")
	if dbUrl == "" {
		log.Fatal("DB Url not found in environment")
	}

	dbConn, err := sql.Open("postgres", dbUrl)
	if err != nil {
		log.Fatal("Failed to connect to databse: ", err)
	}

	apiCfg := apiConfig{
		DB: database.New(dbConn),
	}

	router := chi.NewRouter()

	router.Use(
		cors.Handler(
			cors.Options{
				AllowedOrigins:   []string{"http://*", "https://*"},
				AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
				AllowedHeaders:   []string{"*"},
				ExposedHeaders:   []string{"Link"},
				AllowCredentials: false,
				MaxAge:           300,
			},
		),
	)

	v1Router := chi.NewRouter()
	v1Router.Get("/healthz", handlerReadinessChecker)
	v1Router.Get("/err", handlerErrorChecker)
	v1Router.Post("/users", apiCfg.handlerCreateUser)
	router.Mount("/v1", v1Router)

	server := &http.Server{
		Handler: router,
		Addr:    ":" + port,
	}

	log.Printf("Server starting on port %v", port)
	err = server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
