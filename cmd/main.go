package main

import (
	"fmt"

	"github.com/AlexanderObolonkov/weather-app/configs"
	"github.com/AlexanderObolonkov/weather-app/internal/formatter"
	"github.com/AlexanderObolonkov/weather-app/internal/utils"
	"github.com/AlexanderObolonkov/weather-app/internal/weather"
)

func main() {
	key, err := configs.GetWeatherAPIKey()
	if err != nil {
		utils.ExitWithError(err)
	}
	weatherAPIProvider := weather.NewWeatherAPIProvider(key)

	weatherData, err := weather.GetWeather(weatherAPIProvider)
	if err != nil {
		utils.ExitWithError(err)
	}

	weatherStructured, err := weather.StructWeather(weatherData)
	if err != nil {
		utils.ExitWithError(err)
	}
	formattedWeather := formatter.FormatWeather(weatherStructured)
	fmt.Println(formattedWeather)
}
