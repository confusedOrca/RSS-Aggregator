package server_setup

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() (map[string]string, error) {
	err := godotenv.Load(".env")
	if err != nil {
		return nil, errors.New("failed to load .env file")
	}

	envVariables := map[string]string{
		"PORT":   os.Getenv("PORT"),
		"DB_URL": os.Getenv("DB_URL"),
	}

	if envVariables["PORT"] == "" {
		return nil, errors.New("port not found in environment")
	}

	if envVariables["DB_URL"] == "" {
		return nil, errors.New("DB URL not found in environment")
	}

	return envVariables, nil
}
