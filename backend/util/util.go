package util

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/johneliud/real-time-forum/backend/logger"
	"github.com/johneliud/real-time-forum/database"
	"golang.org/x/crypto/bcrypt"
)

type contextKey string

const UserIDKey contextKey = "userID"

// GetMessagesHandler retrieves messages from the database and returns them as JSON
func GetMessagesHandler(w http.ResponseWriter, r *http.Request) {
	messages, err := database.GetMessages()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(messages)
}

// GetUserProfileHandler retrieves the user's profile data and returns it as JSON
func GetUserProfileHandler(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value(UserIDKey)
	if userID == nil {
		logger.Error("User ID not found in context")
		http.Error(w, "User ID not found in context", http.StatusUnauthorized)
		return
	}

	user, err := database.GetUserProfile(userID.(string))
	if err != nil {
		logger.Error("Error fetching user profile this place", "err", err)
		http.Error(w, "Error fetching user profile", http.StatusInternalServerError)
		return
	}

	if user == nil {
		logger.Error("User not found")
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	profile := map[string]interface{}{
		"name":         user.FirstName + " " + user.LastName,
		"email":        user.Email,
		"profileImage": user.ProfileImage,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(profile)
}

// HashPassword uses bcrypt library to hash sensitive user data before storage.
func HashPassword(password string) (string, error) {
	if len(strings.TrimSpace(password)) == 0 {
		return "", fmt.Errorf("password must contain atleast 6 characters")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("failed generating password hash: %v", err)
	}
	return string(hashedPassword), nil
}
