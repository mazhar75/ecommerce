package middlewares

import (
	"fmt"
	"net/http"
)

func MiddlwareTest1(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Starting testing middleware-1")
		next.ServeHTTP(w, r)
		fmt.Println("Ending testing midleware-1")
	})
}
