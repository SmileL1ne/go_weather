package main

import (
	"encoding/json"
	"net/http"
	"weatherGo/internal/models"
)

func (app *application) errorResponse(w http.ResponseWriter, req *http.Request, status int, message string) {
	errRs := models.ErrorResponse{
		Code: status,
		Msg:  message,
	}

	if err := app.writeJSON(w, errRs, status); err != nil {
		app.logger.Error(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (app *application) serverErrorResponse(w http.ResponseWriter, req *http.Request, err error) {
	app.logger.Error(err.Error())

	message := "the server encountered a problem and could not process your request"
	app.errorResponse(w, req, http.StatusInternalServerError, message)
}

func (app *application) notFoundResponse(w http.ResponseWriter, req *http.Request) {
	message := "the requested resource could not be found"
	app.errorResponse(w, req, http.StatusNotFound, message)
}

func (app *application) notAvailableResponse(w http.ResponseWriter, req *http.Request) {
	message := "the requested resource is currently not available"
	app.errorResponse(w, req, http.StatusServiceUnavailable, message)
}

func (app *application) badRequestResponse(w http.ResponseWriter, req *http.Request, err error) {
	app.errorResponse(w, req, http.StatusBadRequest, err.Error())
}

func (app *application) writeJSON(w http.ResponseWriter, data any, status int) error {
	js, err := json.Marshal(data)
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(js)

	return nil
}
