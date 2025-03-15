package controller

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/johneliud/real-time-forum/backend/logger"
	"github.com/johneliud/real-time-forum/backend/model"
	"github.com/johneliud/real-time-forum/backend/util"
	"github.com/johneliud/real-time-forum/database"
	"golang.org/x/crypto/bcrypt"
)

type SigninResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
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

	token, err := util.GenerateSessionToken()
	if err != nil {
		respondWithError(w, "Authentication error", http.StatusInternalServerError)
		logger.Error("Authentication error: %v", err)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "auth_token",
		Value:    token,
		Path:     "/",
		MaxAge:   3600 * 24,
		HttpOnly: true,
		Secure:   r.TLS != nil,
		SameSite: http.SameSiteLaxMode,
	})

	logger.Info("User authenticated successfully")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(SigninResponse{
		Success: true,
		Message: "Sign in successful",
	})
}
