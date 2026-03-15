package auth

import (
	"fmt"
	"net/http"
)

func GetBearerToken(headers http.Header) (string, error) {
	authHeader := headers.Get("Authorization")
	if authHeader == "" {
		return "", nil
	}

	if len(authHeader) < 7 || authHeader[:7] != "Bearer " {
		return "", fmt.Errorf("invalid Authorization header")
	}

	return authHeader[7:], nil
}
