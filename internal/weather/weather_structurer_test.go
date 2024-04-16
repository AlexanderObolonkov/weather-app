package weather

import (
	"testing"
)

func TestStructWeather(t *testing.T) {
	testData := []byte(`{
		"location": {
			"name": "New York",
			"region": "NY",
			"country": "United States",
			"localtime": "2022-04-10 18:00"
		},
		"current": {
			"last_updated_epoch": 1649653200,
			"temp_c": 15,
			"condition": {
				"text": "Partly cloudy"
			}
		},
		"forecast": {
			"forecastday": [
				{
					"day": {
						"maxtemp_c": 18,
						"mintemp_c": 12,
						"daily_chance_of_rain": 40,
						"daily_chance_of_snow": 0,
						"condition": {
							"text": "Rain showers"
						}
					},
					"hour": [
						{
							"time_epoch": 1649656800,
							"temp_c": 15,
							"condition": {
								"text": "Partly cloudy"
							},
							"cloud": 50,
							"feelslike_c": 15,
							"will_it_rain": 0,
							"chance_of_rain": 0,
							"will_it_snow": 0,
							"chance_of_snow": 0
						}
					]
				}
			]
		}
	}`)

	expectedWeather := Weather{
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
			LastUpdatedEpoch int64        `json:"last_updated_epoch"`
			TempC            TemperatureC `json:"temp_c"`
			Condition        Condition    `json:"condition"`
		}{
			LastUpdatedEpoch: 1649653200,
			TempC:            15,
			Condition: Condition{
				Text: "Partly cloudy",
			},
		},
		Forecast: struct {
			ForecastDay []struct {
				Day struct {
					MaxTempC          TemperatureC `json:"maxtemp_c"`
					MinTempC          TemperatureC `json:"mintemp_c"`
					DailyChanceOfRain int          `json:"daily_chance_of_rain"`
					DailyChanceOfSnow int          `json:"daily_chance_of_snow"`
					Condition         Condition    `json:"condition"`
				} `json:"day"`
				Astro struct {
					Sunrise string `json:"sunrise"`
					Sunset  string `json:"sunset"`
				} `json:"astro"`
				Hour []Hour `json:"hour"`
			} `json:"forecastday"`
		}{
			ForecastDay: []struct {
				Day struct {
					MaxTempC          TemperatureC `json:"maxtemp_c"`
					MinTempC          TemperatureC `json:"mintemp_c"`
					DailyChanceOfRain int          `json:"daily_chance_of_rain"`
					DailyChanceOfSnow int          `json:"daily_chance_of_snow"`
					Condition         Condition    `json:"condition"`
				} `json:"day"`
				Astro struct {
					Sunrise string `json:"sunrise"`
					Sunset  string `json:"sunset"`
				} `json:"astro"`
				Hour []Hour `json:"hour"`
			}{
				{
					Day: struct {
						MaxTempC          TemperatureC `json:"maxtemp_c"`
						MinTempC          TemperatureC `json:"mintemp_c"`
						DailyChanceOfRain int          `json:"daily_chance_of_rain"`
						DailyChanceOfSnow int          `json:"daily_chance_of_snow"`
						Condition         Condition    `json:"condition"`
					}{
						MaxTempC:          18,
						MinTempC:          12,
						DailyChanceOfRain: 40,
						DailyChanceOfSnow: 0,
						Condition: Condition{
							Text: "Rain showers",
						},
					},
					Hour: []Hour{
						{
							TimeEpoch:    1649656800,
							TempC:        15,
							Condition:    Condition{Text: "Partly cloudy"},
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

	resultWeather, err := StructWeather(testData)

	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if !equalWeather(expectedWeather, resultWeather) {
		t.Fatalf("Expected %v but got %v", expectedWeather, resultWeather)
	}
}

func equalWeather(a, b Weather) bool {
	if a.Location != b.Location {
		return false
	}
	if a.Current != b.Current {
		return false
	}
	if len(a.Forecast.ForecastDay) != len(b.Forecast.ForecastDay) {
		return false
	}
	for i := range a.Forecast.ForecastDay {
		if a.Forecast.ForecastDay[i].Day != b.Forecast.ForecastDay[i].Day {
			return false
		}
		if a.Forecast.ForecastDay[i].Astro != b.Forecast.ForecastDay[i].Astro {
			return false
		}
		if len(a.Forecast.ForecastDay[i].Hour) != len(b.Forecast.ForecastDay[i].Hour) {
			return false
		}
		for j := range a.Forecast.ForecastDay[i].Hour {
			if a.Forecast.ForecastDay[i].Hour[j] != b.Forecast.ForecastDay[i].Hour[j] {
				return false
			}
		}
	}
	return true
}
