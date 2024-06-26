package main

import (
	"log"

	"github.com/confusedOrca/RSS-Aggregator/handler"
	"github.com/confusedOrca/RSS-Aggregator/internal/database"
	"github.com/confusedOrca/RSS-Aggregator/server_setup"
	_ "github.com/lib/pq"
)

func main() {

	env, err := server_setup.LoadEnv()
	if err != nil {
		log.Fatal(err)
	}
	port, dbUrl := env["PORT"], env["DB_URL"]

	dbConn, err := server_setup.ConnectToDatabase(dbUrl)
	if err != nil {
		log.Fatal("Failed to connect to databse: ", err)
	}
	defer dbConn.Close()

	apiCfg := handler.ApiConfig{
		DB: database.New(dbConn),
	}

	router := server_setup.SetupRouter(apiCfg)
	server_setup.StartServer(router, port)
}
