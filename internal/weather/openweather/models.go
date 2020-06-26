package openweather

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/jamethy/project-rising-heat/internal/util/ptr"
)

// Config data for open weather map API
type Config struct {
	Latitude  string        `env:"OPEN_WEATHER_LAT"`
	Longitude string        `env:"OPEN_WEATHER_LON"`
	BaseUrl   string        `env:"OPEN_WEATHER_BASE_URL"`
	APIKey    string        `env:"OPEN_WEATHER_API_KEY"`
	Timeout   time.Duration `env:"OPEN_WEATHER_TIMEOUT"`
}

// IsValid returns true if Config is enough to get weather data
func (c *Config) IsValid() bool {
	_, latErr := strconv.ParseFloat(c.Latitude, 10)
	_, lonErr := strconv.ParseFloat(c.Longitude, 10)
	return c.APIKey != "" && c.BaseUrl != "" && latErr == nil && lonErr == nil
}

// OneCall return struct of https://openweathermap.org/api/one-call-api
type OneCall struct {
	Latitude       *float32 `json:"lat"`
	Longitude      *float32 `json:"lon"`
	Timezone       *string  `json:"timezone"`
	TimezoneOffset *int     `json:"timezone_offset"`
	Current        *struct {
		Dt         *int     `json:"dt"`
		Sunrise    *int     `json:"sunrise"`
		Sunset     *int     `json:"sunset"`
		Temp       *float32 `json:"temp"`
		FeelsLike  *float32 `json:"feels_like"`
		Pressure   *float32 `json:"pressure"`
		Humidity   *float32 `json:"humidity"`
		DewPoint   *float32 `json:"dew_point"`
		Uvi        *float32 `json:"uvi"`
		Clouds     *float32 `json:"clouds"`
		Visibility *int     `json:"visibility"`
		WindSpeed  *float32 `json:"wind_speed"`
		WindDeg    *float32 `json:"wind_deg"`

		Weather []*struct {
			ID          *int    `json:"id"`
			Main        *string `json:"main"`
			Description *string `json:"description"`
			Icon        *string `json:"icon"`
		} `json:"weather"`
	} `json:"current"`
}

// GetSunrise helper method to get time struct of sunrise from struct
func (w *OneCall) GetSunrise() *time.Time {
	if w.Current == nil || w.Current.Sunrise == nil {
		return nil
	}
	t := time.Unix(int64(*w.Current.Sunrise), 0)
	if w.Timezone != nil && w.TimezoneOffset != nil {
		t = t.In(time.FixedZone(*w.Timezone, *w.TimezoneOffset))
	}
	return &t
}

// GetSunset helper method to get time struct of sunset from struct
func (w *OneCall) GetSunset() *time.Time {
	if w.Current == nil || w.Current.Sunset == nil {
		return nil
	}
	t := time.Unix(int64(*w.Current.Sunset), 0)
	if w.Timezone != nil && w.TimezoneOffset != nil {
		t = t.In(time.FixedZone(*w.Timezone, *w.TimezoneOffset))
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
func (w *OneCall) GetRainLevel() *float32 {
	if w.Current == nil || w.Current.Weather == nil || len(w.Current.Weather) == 0 {
		return nil
	}

	weather := w.Current.Weather[0]
	if weather.ID == nil {
		return nil
	}

	switch *weather.ID {
	case 230, 231, 300, 301:
		return ptr.Float32(DrizzleRain)
	case 200, 232, 310, 312, 313, 302, 311, 500:
		return ptr.Float32(LightRain)
	case 201, 314, 501, 520, 521:
		return ptr.Float32(MediumRain)
	case 202, 502, 522:
		return ptr.Float32(HeavyRain)
	case 503, 511, 531:
		return ptr.Float32(ExtremeRain)
	default:
		return ptr.Float32(NoRain)
	}
}

// GetRainLevel helper method to calculate rain level from struct
// https://openweathermap.org/weather-conditions#Weather-Condition-Codes-2
func (w *OneCall) GetWeatherDescription() *string {

	if w.Current == nil || w.Current.Weather == nil {
		return nil
	}

	strs := make([]string, len(w.Current.Weather))

	for _, s := range w.Current.Weather {
		strs = append(strs, fmt.Sprintf("[%d: %s]", *s.ID, *s.Description))
	}

	return ptr.Str(strings.Join(strs, " "))
}
