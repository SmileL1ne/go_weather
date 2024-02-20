package mongo

import (
	"context"
	"errors"
	"weatherGo/internal/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type weatherRepository struct {
	db *mongo.Collection
}

func NewWeatherRepository(db *mongo.Collection) *weatherRepository {
	return &weatherRepository{
		db: db,
	}
}

func (wr *weatherRepository) GetByCity(ctx context.Context, city string) (*models.WeatherInfo, error) {
	var weatherInfo models.WeatherInfo

	filter := bson.M{"city": city}

	err := wr.db.FindOne(ctx, filter).Decode(&weatherInfo)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, models.ErrNotFound
		}
		return nil, err
	}

	return &weatherInfo, nil
}

func (wr *weatherRepository) Update(ctx context.Context, city string, info *models.WeatherInfo) error {
	filter := bson.M{"city": city}
	options := options.Update().SetUpsert(true)

	update := bson.M{
		"$set": bson.M{
			"weather":     info.Weather[0].Main,
			"description": info.Weather[0].Description,
			"temp":        info.Main.Temp,
			"feels_like":  info.Main.FeelsLike,
		},
	}

	_, err := wr.db.UpdateOne(ctx, filter, update, options)
	if err != nil {
		return err
	}

	return nil
}
