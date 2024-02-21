package main

import (
	"io"
	"log/slog"
	"os"
	"testing"
	mongoRepo "weatherGo/internal/repository/mongo"
	openweather "weatherGo/internal/repository/openWeather"
	"weatherGo/pkg/mongoDB"

	"github.com/joho/godotenv"
)

func newTestApplication(t *testing.T) *application {
	err := godotenv.Load("../../.env")
	if err != nil {
		t.Fatalf("could not load .env file: %v", err)
	}

	db, err := mongoDB.OpenConnection(os.Getenv("DB_URI"))
	if err != nil {
		t.Fatalf("could not connect to database: %v", err)
	}

	weatherAPIKey := os.Getenv("WEATHER_API_KEY")
	database := os.Getenv("DB_NAME")
	collection := os.Getenv("DB_COLLECTION")
	weatherRepo := mongoRepo.NewWeatherRepository(db.Database(database).Collection(collection))
	weatherAPI := openweather.NewWeatherAPI(weatherAPIKey)

	return &application{
		logger: slog.New(slog.NewTextHandler(io.Discard, nil)),
		wr:     weatherRepo,
		wa:     weatherAPI,
	}
}
