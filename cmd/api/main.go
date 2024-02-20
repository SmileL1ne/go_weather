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
	"weatherGo/pkg/mongoDB"

	"github.com/joho/godotenv"
)

/*
	TODO:
	- add logging in error handling
	- add tests for both handlers
	- add .env.example
	- hold api key in context and retrieve it from there (middleware that stores in context)
	- make fetching func so handler fetches from api through repository
*/

type application struct {
	wr            repository.Database
	weatherAPIKey string
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

	weatherRepo := mongoRepo.NewWeatherRepository(db)
	weatherAPIKey := os.Getenv("WEATHER_API_KEY")

	app := &application{
		wr:            weatherRepo,
		weatherAPIKey: weatherAPIKey,
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
