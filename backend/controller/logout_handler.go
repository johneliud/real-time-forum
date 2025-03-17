package controller

import (
	"encoding/json"
	"net/http"

	"github.com/johneliud/real-time-forum/backend/logger"
	"github.com/johneliud/real-time-forum/database"
)

type LogoutResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

// LogoutHandler handles user logout by invalidating their session.
func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	logger.Info("Logout request received")

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		logger.Error("Method %q not allowed", r.Method)
		return
	}

	// Check if the user has a session cookie
	cookie, err := r.Cookie("auth_token")
	if err != nil {
		logger.Error("No auth cookie found")
		response := LogoutResponse{
			Success: true,
			Message: "Already logged out",
		}
		json.NewEncoder(w).Encode(response)
		return
	}

	// Invalidate the session by clearing it in the database
	_, err = database.DB.Exec("UPDATE users SET session_token = NULL WHERE session_token = ?", cookie.Value)
	if err != nil {
		logger.Error("Error clearing session from database: %v", err)
	}

	// Clear the cookie in the browser
	http.SetCookie(w, &http.Cookie{
		Name:     "auth_token",
		Value:    "",
		Path:     "/",
		MaxAge:   -1,
		HttpOnly: true,
		Secure:   r.TLS != nil,
		SameSite: http.SameSiteLaxMode,
	})

	response := LogoutResponse{
		Success: true,
		Message: "Logged out successfully",
	}
	logger.Info("User logged out successfully")

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
