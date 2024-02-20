package main

import (
	"encoding/json"
	"log"
	"net/http"
	"weatherGo/internal"
)

func (app *application) weather(w http.ResponseWriter, req *http.Request) {
	var input struct {
		City string `json:"city"`
	}

	if err := internal.ReadJSON(req.Body, &input); err != nil {
		// Handle error
		log.Fatalf("Error decoding json from request: %v", err)
	}

	_, err := app.wr.GetByCity(input.City)
	if err != nil {
		log.Print(err)
		w.WriteHeader(500)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Everything is ok!"))
}

func (app *application) weatherPost(w http.ResponseWriter, req *http.Request) {
	var input struct {
		City string `json:"city"`
	}

	if err := json.NewDecoder(req.Body).Decode(&input); err != nil {
		log.Fatalf("Error decoding json from request: %v", err)
	}

}
