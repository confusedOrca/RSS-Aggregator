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

type feedParams struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

func (apiCfg *ApiConfig) HandlerCreateFeed(responseWriter http.ResponseWriter, request *http.Request, user database.User) {
	params := feedParams{}
	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(responseWriter, 400, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}

	feed, err := apiCfg.DB.CreateFeed(request.Context(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
		Url:       params.URL,
		UserID:    user.ID,
	})

	if err != nil {
		respondWithError(responseWriter, 400, fmt.Sprintf("Could not create feed: %v", err))
		return
	}

	formattedFeed := models.DBFeedToFeed(feed)
	respondWithJSON(responseWriter, 200, formattedFeed)
}

func (apiCfg *ApiConfig) HandlerGetFeeds(responseWriter http.ResponseWriter, request *http.Request) {
	feeds, err := apiCfg.DB.Getfeeds(request.Context())

	if err != nil {
		respondWithError(responseWriter, 400, fmt.Sprintf("Could not get feeds: %v", err))
		return
	}

	formattedFeeds := models.DBFeedsToFeeds(feeds)
	respondWithJSON(responseWriter, 200, formattedFeeds)
}
