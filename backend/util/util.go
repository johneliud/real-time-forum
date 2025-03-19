package util

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/johneliud/real-time-forum/database"
	"golang.org/x/crypto/bcrypt"
)

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
	userID := r.Context().Value("userID").(string)

	userProfile, err := database.GetUserProfile(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(userProfile)
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
