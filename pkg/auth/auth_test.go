package auth

import (
	"encoding/base64"
	"net/http"
	"net/http/httptest"
	"testing"
)

// Mock users map for testing
var mockUsers = map[string]any{
	"username": "password",
}

func TestBasicAuth(t *testing.T) {
	// Create a new request with no authentication header
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response
	rr := httptest.NewRecorder()

	// Mock ReadJSON function to return mockUsers map
	users = mockUsers

	// Create a handler using BasicAuth middleware
	handler := BasicAuth(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	// Serve the HTTP request with the handler
	handler.ServeHTTP(rr, req)

	// Check if status code is Unauthorized (401)
	if rr.Code != http.StatusUnauthorized {
		t.Errorf("Handler returned wrong status code: got %v want %v", rr.Code, http.StatusUnauthorized)
	}

	// Check if WWW-Authenticate header is set
	expectedHeader := `Basic realm="Restricted"`
	if rr.Header().Get("WWW-Authenticate") != expectedHeader {
		t.Errorf("Handler did not set correct WWW-Authenticate header: got %v want %v", rr.Header().Get("WWW-Authenticate"), expectedHeader)
	}

	// Create a new request with valid authentication header
	authString := "Basic " + base64.StdEncoding.EncodeToString([]byte("username:password"))
	req, err = http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Authorization", authString)

	// Reset ResponseRecorder for next test
	rr = httptest.NewRecorder()

	// Serve the HTTP request with the handler again
	handler.ServeHTTP(rr, req)

	// Check if status code is OK (200)
	if rr.Code != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", rr.Code, http.StatusOK)
	}
}
