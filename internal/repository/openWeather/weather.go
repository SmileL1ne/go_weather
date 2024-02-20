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

	weather := models.WeatherInfo{}
	if err := internal.ReadJSON(resp.Body, &weather); err != nil {
		return nil, err
	}

	return &weather, nil
}
