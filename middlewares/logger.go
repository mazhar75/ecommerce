package middlewares

import (
	"fmt"
	"net/http"
	"time"
)

// Logger middleware
func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		duration := time.Since(start)
		fmt.Printf("[%s] %s (%v)\n",
			r.Method,
			r.URL.Path,
			duration,
		)
	})
}
