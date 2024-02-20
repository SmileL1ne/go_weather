package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) NewRouter() *httprouter.Router {
	router := httprouter.New()

	router.HandlerFunc(http.MethodGet, "/weather", app.weather)
	router.HandlerFunc(http.MethodPost, "/weather", app.weatherPost)

	return router
}
