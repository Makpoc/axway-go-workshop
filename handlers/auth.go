package handlers

import (
	"net/http"
)

// WithApiTokenAuth wraps a handler function with authentication logic.
// If the request doesn't have a header X-Api-Token or the value of that header isn't 'apiToken' the middeware will
// return 401. If the tokens match - it will invoke the next handler
func WithApiTokenAuth(apiToken string, next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestToken := r.Header.Get("X-Api-Token")
		if requestToken != apiToken {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		next(w, r)
	})
}
