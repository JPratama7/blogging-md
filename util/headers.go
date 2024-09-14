package util

import (
	"net/http"
	"strings"
)

func ExtractBearer(r *http.Request) (string, error) {
	authHeader := r.Header.Get("Authorization")

	if authHeader == "" {
		return "", nil
	}

	prefix := "Bearer "
	if !strings.HasPrefix(authHeader, prefix) {
		return "", nil
	}

	return strings.TrimPrefix(authHeader, prefix), nil
}
