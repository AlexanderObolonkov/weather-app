package weather

import (
	"errors"
	"fmt"
	"io"
	"net/http"
)

var apiServiceError = errors.New("программа не может получить текущую погоду")

type TemperatureC float64

type WeathersProvider interface {
	GetWeather() ([]byte, error)
}

func GetWeather(provider WeathersProvider) ([]byte, error) {
	return provider.GetWeather()
}

type WeatherAPIProvider struct {
	APIKey string
}

func NewWeatherAPIProvider(apiKey string) *WeatherAPIProvider {
	return &WeatherAPIProvider{APIKey: apiKey}
}

func (p *WeatherAPIProvider) GetWeather() ([]byte, error) {
	url := fmt.Sprintf(
		"https://api.weatherapi.com/v1/forecast.json?q=Kolpino&days=1&lang=ru&key=%s",
		p.APIKey,
	)

	res, err := http.Get(url)
	if err != nil {
		return nil, apiServiceError
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return nil, apiServiceError
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, apiServiceError
	}

	return body, nil
}
