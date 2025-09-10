package auth

import (
	"encoding/json"
	"errors"
	"net/http"

	"github/ecommerce/domain/user"
	"github/ecommerce/infra/postgresql"
)

func mapAppErrorToHTTPStatus(err *postgresql.AppError) int {
	switch err.Code {
	case "EMAIL_EXISTS":
		return http.StatusConflict // 409
	case "USER_NOT_FOUND":
		return http.StatusNotFound // 404
	case "INVALID_PASSWORD":
		return http.StatusUnauthorized // 401
	case "DB_QUERY_FAIL", "DB_INSERT_FAIL", "DB_DELETE_FAIL", "DB_RESULT_FAIL":
		return http.StatusInternalServerError // 500
	default:
		return http.StatusInternalServerError
	}
}

func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	var u user.Users

	// Decode JSON body
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&u); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	// Call service layer to create user
	err := h.Service.CreateUser(u)
	if err != nil {
		// Check if it's an AppError
		var appErr *postgresql.AppError
		if errors.As(err, &appErr) {
			// Map error codes to HTTP status codes
			status := mapAppErrorToHTTPStatus(appErr)
			response := map[string]string{
				"error": appErr.Message,
				"code":  appErr.Code,
			}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(status)
			json.NewEncoder(w).Encode(response)
			return
		}

		// Fallback for generic errors
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	// Success
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Registered successfully",
	})
}
