package auth

import (
	"net/http"
	"errors"
	"strings"
)

// GetAPIKey extracts the API key from the header of a HTTP request
func GetAPIKey(headers http.Header) (string, error) {
	apiKey := headers.Get("Authorization")
	if apiKey == "" {
		return "", errors.New("No authorization info found")
	}

	apiKeys := strings.Split(apiKey	, " ")
	if len(apiKeys) != 2 {
		return "", errors.New("Invalid authorization header")
	}

	if apiKeys[0] != "ApiKey" {
		return "", errors.New("Invalid authorization type")
	}

	return apiKeys[1], nil
}