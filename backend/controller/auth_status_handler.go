package controller

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/johneliud/real-time-forum/backend/logger"
	"github.com/johneliud/real-time-forum/database"
)

type AuthStatusResponse struct {
	Authenticated bool   `json:"authenticated"`
	UserID        int64  `json:"userId,omitempty"`
	Message       string `json:"message,omitempty"`
}

// AuthStatusHandler checks if a user is authenticated via session.
func AuthStatusHandler(w http.ResponseWriter, r *http.Request) {
	logger.Info("Auth status check request received")

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		logger.Error("Method %q not allowed", r.Method)
		return
	}

	// Check for auth cookie
	cookie, err := r.Cookie("auth_token")
	if err != nil {
		logger.Error("No auth cookie found")
		response := AuthStatusResponse{
			Authenticated: false,
			Message:       "Not authenticated",
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
		return
	}

	// Validate session in the database
	var userID int64
	err = database.DB.QueryRow("SELECT id FROM users WHERE session_token = ?", cookie.Value).Scan(&userID)
	if err == sql.ErrNoRows {
		logger.Error("Invalid session token")
		response := AuthStatusResponse{
			Authenticated: false,
			Message:       "Session expired",
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
		return
	} else if err != nil {
		logger.Error("Database error checking session: %v", err)
		response := AuthStatusResponse{
			Authenticated: false,
			Message:       "Error checking authentication",
		}
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(response)
		return
	}

	// User is authenticated
	response := AuthStatusResponse{
		Authenticated: true,
		UserID:        userID,
		Message:       "Authenticated",
	}
	logger.Info("User %d is authenticated", userID)
	
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
