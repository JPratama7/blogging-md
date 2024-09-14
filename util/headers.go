package util

import (
	"errors"
	"net/http"
	"strings"
)

func ExtractBearer(r *http.Request) (string, error) {
	authHeader := r.Header.Get("Authorization")

	if authHeader == "" {
		return "", errors.New("Authorization header is required")
	}

	prefix := "Bearer "
	if !strings.HasPrefix(authHeader, prefix) {
		return "", errors.New("Authorization header is invalid")
	}

	return strings.TrimPrefix(authHeader, prefix), nil
}
