package weather

import (
	"errors"
	"testing"
)

type MockProvider struct{}

func (m *MockProvider) GetWeather() (string, error) {
	return "mock weather data", nil
}

func TestGetWeather(t *testing.T) {
	mockProvider := &MockProvider{}

	weatherData, err := GetWeather(mockProvider)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	expected := "mock weather data"
	if weatherData != expected {
		t.Errorf("unexpected weather data, got: %s, want: %s", weatherData, expected)
	}
}

func TestOpenWeatherProvider_GetWeather(t *testing.T) {
	provider := NewOpenWeatherProvider("invalid_api_key")
	_, err := provider.GetWeather()

	if err == nil {
		t.Error("expected an error, got nil")
	}
	if !errors.Is(err, apiServiceError) {
		t.Errorf("expected error %v, got %v", apiServiceError, err)
	}
}