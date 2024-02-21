package main

import (
	"context"
	"fmt"
	"log"
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
	TODO:
	- add logging in error handling
	- add tests for both handlers
	- add .env.example

	2 ways of adding new weather info to database:
	1. current (with put handler - would add if not exist)
	2. add in get handler when not found
*/

type application struct {
	wr repository.Database
	ow repository.WeatherAPI
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	uri := os.Getenv("DB_URI")

	db, err := mongoDB.OpenConnection(uri)
	if err != nil {
		log.Fatalf("Error connection to database: %v", err)
	}
	defer func(ctx context.Context) {
		if err := db.Disconnect(ctx); err != nil {
			return
		}
	}(context.TODO())

	weatherAPIKey := os.Getenv("WEATHER_API_KEY")
	database := os.Getenv("DB_NAME")
	collection := os.Getenv("DB_COLLECTION")
	weatherRepo := mongoRepo.NewWeatherRepository(db.Database(database).Collection(collection))
	weatherAPI := openweather.NewWeatherAPI(weatherAPIKey)

	app := &application{
		wr: weatherRepo,
		ow: weatherAPI,
	}

	srv := http.Server{
		Addr:         fmt.Sprintf("127.0.0.1:%s", os.Getenv("APP_PORT")),
		Handler:      app.NewRouter(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	log.Printf("starting server on %s", srv.Addr)
	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
