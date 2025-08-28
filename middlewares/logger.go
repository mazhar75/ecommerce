package middlewares

import (
	"fmt"
	"net/http"
	"time"
)

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		fmt.Println(("Logger Starting"))
		next.ServeHTTP(w, r)
		df := time.Since(start)
		fmt.Println((df))
		fmt.Println("Logger Ending...")
	})
}
