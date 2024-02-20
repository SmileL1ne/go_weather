package repository

import "weatherGo/internal/models"

type Database interface {
	GetByCity(string) (*models.WeatherInfo, error)
	UpdateCity(string, models.WeatherInfo) error
}

type WeatherAPI interface {
	Fetch(string) (*models.WeatherInfo, error)
}
