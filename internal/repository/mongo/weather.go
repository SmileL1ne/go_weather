package mongo

import (
	"weatherGo/internal/models"

	"go.mongodb.org/mongo-driver/mongo"
)

type weatherRepository struct {
	db *mongo.Client
}

func NewWeatherRepository(db *mongo.Client) *weatherRepository {
	return &weatherRepository{
		db: db,
	}
}

func (wr *weatherRepository) GetByCity(city string) (*models.WeatherInfo, error) {

	return nil, nil
}

func (wr *weatherRepository) UpdateCity(city string, info models.WeatherInfo) error {
	return nil
}
