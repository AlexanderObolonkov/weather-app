package configs

import (
	"errors"
	"os"
)

var errNotSet = errors.New("program can't get api key")

func GetOpenWeatherAPIKey() (string, error) {
	APIKey, err := getEnv("API_KEY")
	if err != nil {
		return "", err
	}
	return APIKey, nil
}

func getEnv(key string) (string, error) {
	if value, exists := os.LookupEnv(key); exists {
		return value, nil
	}
	return "", errNotSet
}
