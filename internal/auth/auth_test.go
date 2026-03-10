package auth

import (
	"errors"
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	authHeader := http.Header{}
	authHeader.Add("Authorization", "ApiKey xxxxxx-1099339")

	_, err := GetAPIKey(authHeader)
	if err != nil {
		t.Errorf("ApiKey expected but not found")
	}
}

func TestReturnsErrNoAuthHeaderIncluded(t *testing.T) {
	authHeader := http.Header{}

	_, err := GetAPIKey(authHeader)
	if err == nil {
		t.Errorf("Expected error: %v", ErrNoAuthHeaderIncluded)
	}
}

func TestErrMalformedAuthorizationHeader(t *testing.T) {
	authHeader := http.Header{}
	authHeader.Add("Authorization", "Bearer the-quick-brown-fox")

	_, err := GetAPIKey(authHeader)
	if err == nil {
		t.Errorf("Expected error: %v", errors.New("malformed authorization header"))
	}
}
