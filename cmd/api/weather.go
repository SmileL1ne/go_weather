package main

import (
	"context"
	"errors"
	"net/http"
	"weatherGo/internal/models"
)

func (app *application) weather(w http.ResponseWriter, req *http.Request) {
	city := req.URL.Query().Get("city")
	if city == "" {
		app.badRequestResponse(w, req, errors.New("city cannot be blank"))
		return
	}

	var info *models.WeatherInfo
	var err error

	info, err = app.wr.GetByCity(context.Background(), city)
	if err != nil && !errors.Is(err, models.ErrNotFound) {
		app.serverErrorResponse(w, req, err)
		return
	}
	if errors.Is(err, models.ErrNotFound) {
		info, err = app.wa.Fetch(city)
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
	city := req.URL.Query().Get("city")
	if city == "" {
		app.badRequestResponse(w, req, errors.New("city cannot be blank"))
		return
	}

	info, err := app.wa.Fetch(city)
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

	id, err := app.wr.Update(context.Background(), city, info)
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
