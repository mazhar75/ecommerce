package middlewares

import (
	"github/ecommerce/drivers"
	"net/http"
	"strings"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		jwt_string := r.Header.Get("Authorization")
		jwt := strings.Split(jwt_string, " ")
		if len(jwt) <= 1 || jwt[1] == "" {
			http.Error(w, "Unauthorized: missing token", http.StatusUnauthorized)
			return
		}
		verification_sucess, err := drivers.ValidateJWT(jwt[1])
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		} else if !verification_sucess {
			http.Error(w, "Invalid JWT", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)

	})
}
