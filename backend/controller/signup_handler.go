package controller

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/johneliud/real-time-forum/backend/logger"
	"github.com/johneliud/real-time-forum/backend/model"
	"github.com/johneliud/real-time-forum/backend/util"
	"github.com/johneliud/real-time-forum/database"
)

type SignupResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

func SignupHandler(w http.ResponseWriter, r *http.Request) {
	logger.Info("Received signup request")

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
		logger.Error("Invalid HTTP method", r.Method)
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	var signupRequest model.SignupRequest
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&signupRequest); err != nil {
		logger.Error("Invalid request body", err)
		respondWithError(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	// Validate input
	if err := util.ValidateSignupRequest(signupRequest); err != nil {
		logger.Error("Failed validating request", err)
		respondWithError(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Check if user with the same email already exists
	var count int
	err := database.DB.QueryRow("SELECT COUNT(*) FROM users WHERE email = ?", signupRequest.Email).Scan(&count)
	if err != nil {
		logger.Error("Database error checking email", err)
		respondWithError(w, "Database error", http.StatusInternalServerError)
		return
	}
	if count > 0 {
		logger.Warn("Email already registered", signupRequest.Email)
		respondWithError(w, "Email already registered", http.StatusBadRequest)
		return
	}

	// Check if user with the same nickname already exists
	err = database.DB.QueryRow("SELECT COUNT(*) FROM users WHERE nick_name = ?", signupRequest.NickName).Scan(&count)
	if err != nil {
		logger.Error("Database error checking nick name", err)
		respondWithError(w, "Database error", http.StatusInternalServerError)
		return
	}
	if count > 0 {
		logger.Warn("Nickname already taken", signupRequest.NickName)
		respondWithError(w, "Nickname already taken", http.StatusBadRequest)
		return
	}

	// Hash password
	hashedPassword, err := util.HashPassword(signupRequest.Password)
	if err != nil {
		logger.Error("Password hashing failed", err)
		respondWithError(w, "Password hashing failed", http.StatusInternalServerError)
		return
	}

	// Insert user into database
	stmt, err := database.DB.Prepare(`
        INSERT INTO users (
            first_name, last_name, nick_name, 
            gender, age, email, password
        ) VALUES (?, ?, ?, ?, ?, ?, ?)
    `)
	if err != nil {
		logger.Error("Database preparation failed", err)
		respondWithError(w, "Database preparation failed", http.StatusInternalServerError)
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(
		signupRequest.FirstName,
		signupRequest.LastName,
		signupRequest.NickName,
		signupRequest.Gender,
		signupRequest.Age,
		signupRequest.Email,
		hashedPassword,
	)
	if err != nil {
		logger.Error("User registration failed", err)
		if strings.Contains(err.Error(), "UNIQUE constraint") {
			if strings.Contains(err.Error(), "users.email") {
				respondWithError(w, "Email already registered", http.StatusBadRequest)
			} else if strings.Contains(err.Error(), "users.nick_name") {
				respondWithError(w, "Nickname already taken", http.StatusBadRequest)
			} else {
				respondWithError(w, "User already exists", http.StatusBadRequest)
			}
			return
		}
		respondWithError(w, "User registration failed", http.StatusBadRequest)
		return
	}

	logger.Info("User registered successfully")
	respondWithSuccess(w, "User registered successfully")
}

func respondWithError(w http.ResponseWriter, message string, status int) {
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(SignupResponse{
		Success: false,
		Message: message,
	})
}

func respondWithSuccess(w http.ResponseWriter, message string) {
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(SignupResponse{
		Success: true,
		Message: message,
	})
}
