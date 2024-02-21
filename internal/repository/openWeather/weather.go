package openweather

import (
	"fmt"
	"net/http"
	"weatherGo/internal"
	"weatherGo/internal/models"
)

type weatherAPI struct {
	weatherAPIKey string
}

func NewWeatherAPI(key string) *weatherAPI {
	return &weatherAPI{
		weatherAPIKey: key,
	}
}

func (a *weatherAPI) Fetch(city string) (*models.WeatherInfo, error) {
	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s&APPID=%s", city, a.weatherAPIKey)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		if resp.StatusCode == http.StatusNotFound {
			return nil, models.ErrNotFound
		}
		return nil, models.ErrNotAvailable
	}

	raw := models.WeatherInfoRaw{}
	if err := internal.ReadJSON(resp.Body, &raw); err != nil {
		return nil, err
	}

	info := models.WeatherInfo{
		Weather:     raw.Weather[0].Main,
		Description: raw.Weather[0].Description,
		Temp:        internal.ConvertKelvinToCelsius(raw.Main.Temp),
		FeelsLike:   internal.ConvertKelvinToCelsius(raw.Main.FeelsLike),
	}

	return &info, nil
}
