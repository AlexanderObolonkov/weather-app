package formatter

import (
	"testing"

	"github.com/AlexanderObolonkov/weather-app/internal/weather"
)

func TestFormatWeather(t *testing.T) {
	testWeather := weather.Weather{
		Location: struct {
			Name      string `json:"name"`
			Region    string `json:"region"`
			Country   string `json:"country"`
			LocalTime string `json:"localtime"`
		}{
			Name:      "New York",
			Region:    "NY",
			Country:   "United States",
			LocalTime: "2022-04-10 18:00",
		},
		Current: struct {
			LastUpdatedEpoch int64                `json:"last_updated_epoch"`
			TempC            weather.TemperatureC `json:"temp_c"`
			Condition        weather.Condition    `json:"condition"`
		}{
			LastUpdatedEpoch: 1649653200,
			TempC:            15,
			Condition: weather.Condition{
				Text: "Partly cloudy",
			},
		},
		Forecast: struct {
			ForecastDay []struct {
				Day struct {
					MaxTempC          weather.TemperatureC `json:"maxtemp_c"`
					MinTempC          weather.TemperatureC `json:"mintemp_c"`
					DailyChanceOfRain int                  `json:"daily_chance_of_rain"`
					DailyChanceOfSnow int                  `json:"daily_chance_of_snow"`
					Condition         weather.Condition    `json:"condition"`
				} `json:"day"`
				Astro struct {
					Sunrise string `json:"sunrise"`
					Sunset  string `json:"sunset"`
				} `json:"astro"`
				Hour []weather.Hour `json:"hour"`
			} `json:"forecastday"`
		}{
			ForecastDay: []struct {
				Day struct {
					MaxTempC          weather.TemperatureC `json:"maxtemp_c"`
					MinTempC          weather.TemperatureC `json:"mintemp_c"`
					DailyChanceOfRain int                  `json:"daily_chance_of_rain"`
					DailyChanceOfSnow int                  `json:"daily_chance_of_snow"`
					Condition         weather.Condition    `json:"condition"`
				} `json:"day"`
				Astro struct {
					Sunrise string `json:"sunrise"`
					Sunset  string `json:"sunset"`
				} `json:"astro"`
				Hour []weather.Hour `json:"hour"`
			}{
				{
					Day: struct {
						MaxTempC          weather.TemperatureC `json:"maxtemp_c"`
						MinTempC          weather.TemperatureC `json:"mintemp_c"`
						DailyChanceOfRain int                  `json:"daily_chance_of_rain"`
						DailyChanceOfSnow int                  `json:"daily_chance_of_snow"`
						Condition         weather.Condition    `json:"condition"`
					}{
						MaxTempC:          18,
						MinTempC:          12,
						DailyChanceOfRain: 40,
						DailyChanceOfSnow: 0,
						Condition: weather.Condition{
							Text: "Rain showers",
						},
					},
					Astro: struct {
						Sunrise string `json:"sunrise"`
						Sunset  string `json:"sunset"`
					}{
						Sunrise: "06:30",
						Sunset:  "19:00",
					},
					Hour: []weather.Hour{
						{
							TimeEpoch:    1649656800,
							TempC:        15,
							Condition:    weather.Condition{Text: "Partly cloudy"},
							Cloud:        50,
							FeelsLikeC:   15,
							WillItRain:   0,
							ChanceOfRain: 0,
							WillItSnow:   0,
							ChanceOfSnow: 0,
						},
					},
				},
			},
		},
	}

	expectedOutput := `Локация: New York, NY, United States
Дата и время: 2022-04-10 18:00
На момент 05:00: 15.0°C, partly cloudy
Восход: 06:30, закат: 19:00

В 06:00 15.0°C, ощущается как 15.0°C, partly cloudy
Облачность 50%, вероятность дождя: 0%, снега: 0%
`

	result := FormatWeather(testWeather)

	if result != expectedOutput {
		t.Errorf("Unexpected output:\nExpected: %s\nGot: %s", expectedOutput, result)
	}
}
