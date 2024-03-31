package main

import (
	"fmt"
	"os"

	"github.com/AlexanderObolonkov/weather-app/configs"
	"github.com/AlexanderObolonkov/weather-app/internal/weather"
)

func main() {
	key, err := configs.GetOpenWeatherAPIKey()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	weatherAPIProvider := weather.NewWeatherAPIProvider(key)

	weatherData, err := weather.GetWeather(weatherAPIProvider)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	weatherFormatted, err := weather.FormatWeather(weatherData)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(weatherFormatted)
}
