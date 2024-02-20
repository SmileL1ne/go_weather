package mongo

import "go.mongodb.org/mongo-driver/mongo"

type WeatherRepository struct {
	db *mongo.Client
}

func NewWeatherRepository(db *mongo.Client) *WeatherRepository {
	return &WeatherRepository{
		db: db,
	}
}
