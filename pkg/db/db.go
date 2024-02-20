package db

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Add logs to error handling

func OpenConnection(uri string) (*mongo.Client, error) {
	dbCtx, dbCancelCtx := context.WithTimeout(context.Background(), 100*time.Second)
	defer dbCancelCtx()

	options := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(dbCtx, options)
	if err != nil {
		return nil, err
	}

	if err := client.Ping(context.Background(), nil); err != nil {
		if err := client.Disconnect(context.Background()); err != nil {
			panic(err)
		}
		return nil, err
	}

	return client, err
}
