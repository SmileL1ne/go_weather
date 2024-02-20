package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func (app *application) weather(w http.ResponseWriter, req *http.Request) {
	city := "Astana"

	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s&APPID=%s", city, app.weatherAPIKey)

	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("Error making GET request to Open Weather API: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Unexpected response status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading response body: %v", err)
	}

	var data map[string]interface{}
	if err := json.Unmarshal(body, &data); err != nil {
		log.Fatalf("Error unmarshaling response body: %v", err)
	}
	fmt.Println(data)

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
