package weather

import (
	"errors"
	"fmt"
	"io"
	"net/http"
)

var apiServiceError = errors.New("program can't current weather")

type WeathersProvider interface {
	GetWeather() (string, error)
}

type OpenWeatherProvider struct {
	APIKey string
}

func NewOpenWeatherProvider(apiKey string) *OpenWeatherProvider {
	return &OpenWeatherProvider{APIKey: apiKey}
}

func (p *OpenWeatherProvider) GetWeather() (string, error) {
	url := fmt.Sprintf(
		"http://api.openweathermap.org/data/2.5/weather?q=Kolpino&appid=%s&lang=ru&units=metric",
		p.APIKey,
	)

	res, err := http.Get(url)
	if err != nil {
		return "", apiServiceError
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return "", apiServiceError
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", apiServiceError
	}

	return string(body), nil
}

func GetWeather(provider WeathersProvider) (string, error) {
	return provider.GetWeather()
}
