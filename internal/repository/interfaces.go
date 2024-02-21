package repository

import (
	"context"
	"weatherGo/internal/models"
)

type Database interface {
	GetByCity(context.Context, string) (*models.WeatherInfo, error)
	Update(context.Context, string, *models.WeatherInfo) (string, error)
}

type WeatherAPI interface {
	Fetch(string) (*models.WeatherInfo, error)
}
