package main

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"time"
	"weatherGo/internal/repository"
	mongoRepo "weatherGo/internal/repository/mongo"
	openweather "weatherGo/internal/repository/openWeather"
	"weatherGo/pkg/mongoDB"

	"github.com/joho/godotenv"
)

/*
	2 ways of adding new weather info to database:

	1. current (with put handler - would add if not exist)
	2. add in get handler when not found
*/

type application struct {
	wr     repository.Database
	wa     repository.WeatherAPI
	logger *slog.Logger
}

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	err := godotenv.Load()
	if err != nil {
		logger.Error("could not load .env file", "error", err)
		os.Exit(1)
	}

	db, err := mongoDB.OpenConnection(os.Getenv("DB_URI"))
	if err != nil {
		logger.Error("could not connect to database", "error", err)
		os.Exit(1)
	}
	defer func(ctx context.Context) {
		if err := db.Disconnect(ctx); err != nil {
			return
		}
	}(context.TODO())

	logger.Info("sucessfully connected to database")

	weatherAPIKey := os.Getenv("WEATHER_API_KEY")
	database := os.Getenv("DB_NAME")
	collection := os.Getenv("DB_COLLECTION")
	weatherRepo := mongoRepo.NewWeatherRepository(db.Database(database).Collection(collection))
	weatherAPI := openweather.NewWeatherAPI(weatherAPIKey)

	app := &application{
		wr:     weatherRepo,
		wa:     weatherAPI,
		logger: logger,
	}

	srv := http.Server{
		Addr:         fmt.Sprintf("0.0.0.0:%s", os.Getenv("APP_PORT")),
		Handler:      app.NewRouter(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	logger.Info("starting server", "addr", srv.Addr)
	if err := srv.ListenAndServe(); err != nil {
		logger.Error("ListenAndServe error", "error", err)
		os.Exit(1)
	}
}
