package auth

import (
	"errors"
	"net/http"
	"strings"
)

// expected header format:
// Authorization: APIKey <insert api key here>
func GetAPIKey(headers http.Header) (string, error) {
	authVal := headers.Get("Authorization")
	if authVal == "" {
		return "", errors.New("no authentication info found")
	}

	values := strings.Split(authVal, " ")
	if len(values) != 2 || values[0] != "APIKey" {
		return "", errors.New("invalid authorization header")
	}

	return values[1], nil
}
