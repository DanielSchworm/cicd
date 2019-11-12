package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHelloWorldHandler(t *testing.T) {
	t.Run("test1", func(t *testing.T) {
		req, err := http.NewRequest("GET", "localhost:8081", nil)
		if err != nil {
			t.Fatalf("could not created request: %v", err)
		}

		rec := httptest.NewRecorder()
		handler := http.HandlerFunc(printHelloWorld)

		handler.ServeHTTP(rec, req)

		res := rec.Result()
		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			t.Errorf("expected status OK, got %v", res.Status)
		}
	})

	t.Run("test2", func(t *testing.T) {
		req, err := http.NewRequest("GET", "localhost:8081", nil)
		if err != nil {
			t.Fatalf("could not created request: %v", err)
		}
		rec := httptest.NewRecorder()
		handler := http.HandlerFunc(printHelloWorld)

		handler.ServeHTTP(rec, req)

		expected := "Hello World"
		if rec.Body.String() != expected {
			t.Errorf("handler returned unexpected body: got %v want %v", rec.Body.String(), expected)
		}
	})
}
