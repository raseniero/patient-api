package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestGreetHandler(t *testing.T) {
	// Test cases for the greet handler
	testCases := []struct {
		path     string
		expected string
	}{
		{"/", "Hello, Gopher!"},
		{"/John", "Hello, John!"},
		{"/Alice", "Hello, Alice!"},
	}

	for _, tc := range testCases {
		req, err := http.NewRequest("GET", tc.path, nil)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(greet)

		handler.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusOK)
		}

		expected := tc.expected
		if !strings.Contains(rr.Body.String(), expected) {
			t.Errorf("handler returned unexpected body: got %v want %v",
				rr.Body.String(), expected)
		}
	}
}

func TestVersionHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/version", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(version)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check if the response contains HTML and build info
	if !strings.Contains(rr.Body.String(), "<!DOCTYPE html>") {
		t.Error("handler did not return HTML content")
	}
}
