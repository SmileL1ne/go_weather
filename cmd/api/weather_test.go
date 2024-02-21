package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"weatherGo/internal/assert"
	"weatherGo/internal/models"
)

func TestWeatherSuccess(t *testing.T) {
	app := newTestApplication(t)

	ts := httptest.NewServer(app.NewRouter())
	defer ts.Close()

	rs, err := ts.Client().Get(ts.URL + "/weather?city=Astana")
	if err != nil {
		t.Fatal(err)
	}
	defer rs.Body.Close()

	assert.Equal(t, rs.StatusCode, http.StatusOK)

	var rsData map[string]interface{}
	if err := json.NewDecoder(rs.Body).Decode(&rsData); err != nil {
		t.Fatal(err)
	}

	fields := []string{"weather", "description", "temp", "feels_like"}
	for _, field := range fields {
		if _, ok := rsData[field]; !ok {
			t.Errorf("request body is missing field: %s", field)
		}
	}
}

func TestWeatherNotFound(t *testing.T) {
	app := newTestApplication(t)

	ts := httptest.NewServer(app.NewRouter())
	defer ts.Close()

	rs, err := ts.Client().Get(ts.URL + "/weather?city=asdfasdf")
	if err != nil {
		t.Fatal(err)
	}
	defer rs.Body.Close()

	assert.Equal(t, rs.StatusCode, http.StatusNotFound)

	var rsData models.ErrorResponse
	if err := json.NewDecoder(rs.Body).Decode(&rsData); err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, rsData.Msg, "the requested resource could not be found")
}
