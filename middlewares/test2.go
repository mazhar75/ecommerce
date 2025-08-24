package middlewares

import (
	"fmt"
	"net/http"
)

func MiddlewareTest2(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Starting middleware testing-2")
		next.ServeHTTP(w, r)
		fmt.Println("Ending middleware testing-2")
	})
}
