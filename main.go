package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
	"weatherGo/pkg/db"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
)

/*
	TODO:
	- add logging in error handling
	- add tests for both handlers
	- add .env.example
*/

type application struct {
	db *mongo.Client
}

func main() {
	app := application{}

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	uri := os.Getenv("DB_URI")

	db, err := db.OpenConnection(uri)
	if err != nil {
		log.Fatalf("Error connection to database: %v", err)
	}
	defer func(ctx context.Context) {
		if err := db.Disconnect(ctx); err != nil {
			return
		}
	}(context.TODO())

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
