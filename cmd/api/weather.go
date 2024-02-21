package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"weatherGo/internal"
	"weatherGo/internal/models"
)

func (app *application) weather(w http.ResponseWriter, req *http.Request) {
	var input struct {
		City string `json:"city"`
	}

	if err := internal.ReadJSON(req.Body, &input); err != nil {
		app.serverErrorResponse(w, req, err)
		return
	}

	var info *models.WeatherInfo
	var err error

	info, err = app.wr.GetByCity(context.Background(), input.City)
	if err != nil && !errors.Is(err, models.ErrNotFound) {
		app.serverErrorResponse(w, req, err)
		return
	}
	if errors.Is(err, models.ErrNotFound) {
		info, err = app.ow.Fetch(input.City)
		if err != nil {
			switch {
			case errors.Is(err, models.ErrNotFound):
				app.notFoundResponse(w, req)
			case errors.Is(err, models.ErrNotAvailable):
				app.notAvailableResponse(w, req)
			default:
				app.serverErrorResponse(w, req, err)
			}
			return
		}
	}

	if err := app.writeJSON(w, info, http.StatusOK); err != nil {
		app.serverErrorResponse(w, req, err)
		return
	}
}

func (app *application) weatherPost(w http.ResponseWriter, req *http.Request) {
	var input struct {
		City string `json:"city"`
	}

	if err := internal.ReadJSON(req.Body, &input); err != nil {
		log.Fatalf("Error decoding json from request: %v", err)
	}

	info, err := app.ow.Fetch(input.City)
	if err != nil {
		switch {
		case errors.Is(err, models.ErrNotFound):
			app.notFoundResponse(w, req)
		case errors.Is(err, models.ErrNotAvailable):
			app.notAvailableResponse(w, req)
		default:
			app.serverErrorResponse(w, req, err)
		}
		return
	}

	id, err := app.wr.Update(context.Background(), input.City, info)
	if err != nil {
		app.serverErrorResponse(w, req, err)
		return
	}

	success := map[string]string{
		"status": "success",
	}
	if id != "" {
		success["inserted_id"] = id
	}

	if err := app.writeJSON(w, success, http.StatusOK); err != nil {
		app.serverErrorResponse(w, req, err)
		return
	}
}
