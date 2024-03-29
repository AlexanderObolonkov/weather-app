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
	openWeatherProvider := weather.NewOpenWeatherProvider(key)

	weatherData, err := weather.GetWeather(openWeatherProvider)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(weatherData)
}
