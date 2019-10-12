// Package ratelmt provides rate limiter middleware
package ratelmt

import (
	"net/http"

	"golang.org/x/time/rate"
)

// Mw rate limits http requests to r reqs/s
func Mw(r float64, next http.Handler) http.HandlerFunc {
	l := rate.NewLimiter(rate.Limit(r), int(r))
	return func(w http.ResponseWriter, r *http.Request) {
		if !l.Allow() {
			http.Error(w, http.StatusText(429), 429)
			return
		}
		next.ServeHTTP(w, r)
	}
}
