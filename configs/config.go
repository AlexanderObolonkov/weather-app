package configs

import (
	"errors"
	"os"
)

var errNotSet = errors.New("программа не может получить ключ api")

func GetWeatherAPIKey() (string, error) {
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
