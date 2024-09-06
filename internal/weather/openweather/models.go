package openweather

import (
	"fmt"
	"strings"
	"time"
)

// Config data for open weather map API
type Config struct {
	Latitude  float64       `json:"latitude"`
	Longitude float64       `json:"longitude"`
	BaseUrl   string        `json:"baseUrl"`
	APIKey    string        `json:"apiKey"`
	Timeout   time.Duration `json:"timeout"`
}

// IsValid returns true if Config is enough to get weather data
func (c *Config) IsValid() bool {
	return c.APIKey != "" && c.BaseUrl != "" &&
		c.Latitude != 0 && c.Longitude != 0 // sorry person who lives on their boat here
}

// OneCall return struct of https://openweathermap.org/api/one-call-api
type OneCall struct {
	Latitude       float64 `json:"lat"`
	Longitude      float64 `json:"lon"`
	Timezone       string  `json:"timezone"`
	TimezoneOffset int     `json:"timezone_offset"`
	Current        struct {
		Dt         int     `json:"dt"`
		Sunrise    int     `json:"sunrise"`
		Sunset     int     `json:"sunset"`
		Temp       float64 `json:"temp"`
		FeelsLike  float64 `json:"feels_like"`
		Pressure   float64 `json:"pressure"`
		Humidity   float64 `json:"humidity"`
		DewPoint   float64 `json:"dew_point"`
		Uvi        float64 `json:"uvi"`
		Clouds     float64 `json:"clouds"`
		Visibility int     `json:"visibility"`
		WindSpeed  float64 `json:"wind_speed"`
		WindDeg    float64 `json:"wind_deg"`

		Weather []struct {
			ID          int    `json:"id"`
			Main        string `json:"main"`
			Description string `json:"description"`
			Icon        string `json:"icon"`
		} `json:"weather"`
	} `json:"current"`
}

// GetSunrise helper method to get time struct of sunrise from struct
func (w *OneCall) GetSunrise() *time.Time {
	if w.Current.Sunrise == 0 {
		return nil
	}
	t := time.Unix(int64(w.Current.Sunrise), 0)
	if w.Timezone != "" && w.TimezoneOffset != 0 {
		t = t.In(time.FixedZone(w.Timezone, w.TimezoneOffset))
	}
	return &t
}

// GetSunset helper method to get time struct of sunset from struct
func (w *OneCall) GetSunset() *time.Time {
	if w.Current.Sunset == 0 {
		return nil
	}
	t := time.Unix(int64(w.Current.Sunset), 0)
	if w.Timezone != "" && w.TimezoneOffset != 0 {
		t = t.In(time.FixedZone(w.Timezone, w.TimezoneOffset))
	}
	return &t
}

const NoRain = 0
const DrizzleRain = 0.5
const LightRain = 1
const MediumRain = 4
const HeavyRain = 7
const ExtremeRain = 10

// GetRainLevel helper method to calculate rain level from struct
// https://openweathermap.org/weather-conditions#Weather-Condition-Codes-2
func (w *OneCall) GetRainLevel() float64 {
	if len(w.Current.Weather) == 0 {
		return -1
	}

	weather := w.Current.Weather[0]
	if weather.ID == 0 {
		return -1
	}

	switch weather.ID {
	case 230, 231, 300, 301:
		return DrizzleRain
	case 200, 232, 310, 312, 313, 302, 311, 500:
		return LightRain
	case 201, 314, 501, 520, 521:
		return MediumRain
	case 202, 502, 522:
		return HeavyRain
	case 503, 511, 531:
		return ExtremeRain
	default:
		return NoRain
	}
}

// GetWeatherDescription helper method to calculate rain level from struct
// https://openweathermap.org/weather-conditions#Weather-Condition-Codes-2
func (w *OneCall) GetWeatherDescription() string {

	strs := make([]string, len(w.Current.Weather))

	for _, s := range w.Current.Weather {
		strs = append(strs, fmt.Sprintf("[%d: %s]", s.ID, s.Description))
	}

	return strings.Join(strs, " ")
}
