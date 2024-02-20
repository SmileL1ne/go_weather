package main

import (
	"context"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type application struct {
}

func main() {
	app := application{}

	router := httprouter.New()
	router.HandlerFunc(http.MethodGet, "/weather", app.weather)
	router.HandlerFunc(http.MethodPost, "/weather", app.weatherPost)
}

func (app *application) weather(w http.ResponseWriter, req *http.Request) {

}

func (app *application) weatherPost(w http.ResponseWriter, req *http.Request) {

}

func OpenConnection(uri string) (*mongo.Client, error) {
	options := options.Client().ApplyURI(uri)

	client, err := mongo.Connect(context.Background(), options)
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
