package formatter

import (
	"fmt"
	"strings"
	"time"

	"github.com/AlexanderObolonkov/weather-app/internal/weather"
)

func FormatWeather(weather weather.Weather) string {
	output := strings.Builder{}
	output.WriteString(formatLocation(weather))
	output.WriteString(formatCurrent(weather))
	output.WriteString(formatAstro(weather))
	for _, hour := range weather.Forecast.ForecastDay[0].Hour {
		output.WriteString("\n")
		output.WriteString(formatHour(hour))
	}
	return output.String()
}

func formatLocation(weather weather.Weather) string {
	return fmt.Sprintf(
		"Локация: %s, %s, %s\nДата и время: %s\n",
		weather.Location.Name,
		weather.Location.Region,
		weather.Location.Country,
		weather.Location.LocalTime,
	)
}

func formatCurrent(weather weather.Weather) string {
	currentTime := time.Unix(weather.Current.LastUpdatedEpoch, 0)
	return fmt.Sprintf(
		"На момент %s: %.1f°C, %s\n",
		currentTime.Format("15:04"),
		weather.Current.TempC,
		strings.ToLower(weather.Current.Condition.Text),
	)
}

func formatAstro(weather weather.Weather) string {
	forecastDay := weather.Forecast.ForecastDay[0]
	return fmt.Sprintf(
		"Восход: %s, закат: %s\n",
		forecastDay.Astro.Sunrise,
		forecastDay.Astro.Sunset,
	)
}

func formatHour(hour weather.Hour) string {
	currentTime := time.Unix(hour.TimeEpoch, 0)
	return fmt.Sprintf(
		"В %s %.1f°C, ощущается как %.1f°C, %s\n"+
			"Облачность %d%%, вероятность дождя: %d%%, снега: %d%%\n",
		currentTime.Format("15:04"),
		hour.TempC,
		hour.FeelsLikeC,
		strings.ToLower(hour.Condition.Text),
		hour.Cloud,
		hour.ChanceOfRain,
		hour.ChanceOfSnow,
	)
}
