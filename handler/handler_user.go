package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/confusedOrca/RSS-Aggregator/internal/database"
	"github.com/confusedOrca/RSS-Aggregator/models"
	"github.com/google/uuid"
)

type parameters struct {
	Name string `json:"name"`
}

func (apiCfg *ApiConfig) HandlerCreateUser(responseWriter http.ResponseWriter, request *http.Request) {
	decoder := json.NewDecoder(request.Body)
	params := parameters{}

	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(responseWriter, 400, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}

	user, err := apiCfg.DB.CreateUser(request.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
	})

	if err != nil {
		respondWithError(responseWriter, 400, fmt.Sprintf("Could not create user: %v", err))
		return
	}

	formatted_user := models.DBUserToUser(user)
	respondWithJSON(responseWriter, 200, formatted_user)
}
