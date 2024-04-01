package weather

import (
	"encoding/json"
	"errors"
)

var formatterError = errors.New("программа не может форматировать данные")

type Weather struct {
	Location struct {
		Name      string `json:"name"`
		Region    string `json:"region"`
		Country   string `json:"country"`
		LocalTime string `json:"localtime"`
	} `json:"location"`
	Current struct {
		LastUpdatedEpoch int64        `json:"last_updated_epoch"`
		TempC            TemperatureC `json:"temp_c"`
		Condition        Condition    `json:"condition"`
	} `json:"current"`
	Forecast struct {
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
	} `json:"forecast"`
}

type Condition struct {
	Text string `json:"text"`
}

type Hour struct {
	TimeEpoch    int64        `json:"time_epoch"`
	TempC        TemperatureC `json:"temp_c"`
	Condition    Condition    `json:"condition"`
	Cloud        int          `json:"cloud"`
	FeelsLikeC   TemperatureC `json:"feelslike_c"`
	WillItRain   int          `json:"will_it_rain"`
	ChanceOfRain int          `json:"chance_of_rain"`
	WillItSnow   int          `json:"will_it_snow"`
	ChanceOfSnow int          `json:"chance_of_snow"`
}

func StructWeather(weatherData []byte) (Weather, error) {
	var weather Weather
	err := json.Unmarshal(weatherData, &weather)
	if err != nil {
		return Weather{}, formatterError
	}
	return weather, nil
}
