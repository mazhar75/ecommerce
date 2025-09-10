package auth

import (
	"encoding/json"
	"errors"
	"net/http"

	"github/ecommerce/domain/user"
	"github/ecommerce/infra/postgresql"
)

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var u user.Users

	// Decode JSON body
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&u); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	// Call service layer to create user
	err := h.Service.Login(u.Email, u.Password)
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
		"message": "Logged in successfully",
	})
}
