package controller

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/johneliud/real-time-forum/backend/logger"
	"github.com/johneliud/real-time-forum/backend/model"
	"github.com/johneliud/real-time-forum/database"
	"golang.org/x/crypto/bcrypt"
)

type SigninResponse struct {
	Success      bool   `json:"success"`
	Message      string `json:"message"`
	SessionToken string `json:"sessionToken"`
}

type SigninRequest struct {
	Identifier string `json:"identifier"`
	Password   string `json:"password"`
}

func SigninHandler(w http.ResponseWriter, r *http.Request) {
	logger.Info("Received signin request")

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == http.MethodOptions {
		logger.Info("OPTIONS request, returning", r.Method)
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		logger.Error("Method %q not allowed", r.Method)
		return
	}

	var signinRequest SigninRequest
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&signinRequest); err != nil {
		respondWithError(w, "Invalid request body", http.StatusBadRequest)
		logger.Error("Invalid request body %v", err)
		return
	}
	defer r.Body.Close()

	if signinRequest.Identifier == "" || signinRequest.Password == "" {
		respondWithError(w, "Email/nickname and password are required", http.StatusBadRequest)
		logger.Error("Email/nickname and password are required")
		return
	}

	var user model.User
	var hashedPassword string

	query := `SELECT id, first_name, last_name, nick_name, gender, age, email, password 
              FROM users 
              WHERE email = ? OR nick_name = ?`

	err := database.DB.QueryRow(query, signinRequest.Identifier, signinRequest.Identifier).Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.NickName,
		&user.Gender,
		&user.Age,
		&user.Email,
		&hashedPassword,
	)

	if err == sql.ErrNoRows {
		respondWithError(w, "Invalid credentials", http.StatusUnauthorized)
		logger.Warn("Invalid credentials: %v", err)
		return
	} else if err != nil {
		respondWithError(w, "Database error", http.StatusInternalServerError)
		logger.Error("Database error: %v", err)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(signinRequest.Password))
	if err != nil {
		respondWithError(w, "Invalid credentials", http.StatusUnauthorized)
		logger.Warn("Invalid credentials: %v", err)
		return
	}

	sessionToken := uuid.New().String()
	sessionDuration := 24 * time.Hour

	// Update the user's session token in the database
	_, err = database.DB.Exec("UPDATE users SET session_token = ? WHERE id = ?", sessionToken, user.ID)
	if err != nil {
		respondWithError(w, "Error creating session", http.StatusInternalServerError)
		logger.Error("Error saving session to database: %v", err)
		return
	}

	// Set the session cookie
	http.SetCookie(w, &http.Cookie{
		Name:     "session_token",
		Value:    sessionToken,
		Path:     "/",
		MaxAge:   int(sessionDuration.Seconds()),
		HttpOnly: true,
		Secure:   r.TLS != nil,
		SameSite: http.SameSiteLaxMode,
	})

	logger.Info("Setting session_token cookie with value: %s", sessionToken)
	logger.Info("User authenticated successfully")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(SigninResponse{
		Success:      true,
		Message:      "Sign in successful",
		SessionToken: sessionToken,
	})
}
