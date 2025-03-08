package controller

import (
	"encoding/json"
	"net/http"

	"github.com/johneliud/real-time-forum/backend/model"
	"github.com/johneliud/real-time-forum/backend/util"
	"github.com/johneliud/real-time-forum/database"
)

type SignupResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

func SignupHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	var signupRequest model.SignupRequest
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&signupRequest); err != nil {
		respondWithError(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	// Validate input
	if err := util.ValidateSignupRequest(signupRequest); err != nil {
		respondWithError(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Check if user with the same email already exists
	var count int
	err := database.DB.QueryRow("SELECT COUNT(*) FROM users WHERE email = ?", signupRequest.Email).Scan(&count)
	if err != nil {
		respondWithError(w, "Database error", http.StatusInternalServerError)
		return
	}
	if count > 0 {
		respondWithError(w, "Email already registered", http.StatusBadRequest)
		return
	}

	// Check if user with the same nickname already exists
	err = database.DB.QueryRow("SELECT COUNT(*) FROM users WHERE nick_name = ?", signupRequest.NickName).Scan(&count)
	if err != nil {
		respondWithError(w, "Database error", http.StatusInternalServerError)
		return
	}
	if count > 0 {
		respondWithError(w, "Nickname already taken", http.StatusBadRequest)
		return
	}

	// Hash password
	hashedPassword, err := util.HashPassword(signupRequest.Password)
	if err != nil {
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
		respondWithError(w, "User registration failed", http.StatusBadRequest)
		return
	}

	// Respond with success
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
