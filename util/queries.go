package util

import (
	"github.com/go-chi/chi/v5"
	"net/http"
)

func ExtractQuery(r *http.Request, key string, value ...string) string {
	query := r.URL.Query()

	if !query.Has(key) {
		if len(value) > 0 {
			return value[0]
		}
		return ""
	}

	return query.Get(key)
}

func ExtractURL(r *http.Request, key string, value ...string) string {
	param := chi.URLParam(r, key)
	if param == "" {
		if len(value) > 0 {
			return value[0]
		}
		return ""
	}

	return param
}
