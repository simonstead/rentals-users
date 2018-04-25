package main

import (
	"encoding/json"
	"github.com/simonstead/rentals-users/handlers"
	"github.com/simonstead/rentals-users/structs"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestIndexRoute(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(handlers.RootHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code, got %v but wanted %v", status, http.StatusOK)
	}

	result := structs.JsonRepsonse{}
	if err := json.NewDecoder(rr.Body).Decode(&result); err != nil {
		t.Errorf("body did not return correct json response: %v", err)
	}

	if result.Msg != "success" {
		t.Errorf("json response returned but it errored:\n %v", result.Error)
	}
}
