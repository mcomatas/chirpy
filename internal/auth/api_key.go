package auth

import (
	"fmt"
	"net/http"
)

func GetAPIKey(headers http.Header) (string, error) {
	authHeader := headers.Get("Authorization")
	if authHeader == "" {
		return "", fmt.Errorf("no authorization header")
	}

	if len(authHeader) < 8 || authHeader[:7] != "ApiKey " {
		return "", fmt.Errorf("invalid Authorization header")
	}

	return authHeader[7:], nil
}
