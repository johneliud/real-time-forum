package controller

import (
	"encoding/json"
	"net/http"

	"github.com/johneliud/real-time-forum/backend/logger"
)

type ErrorResponse struct {
	StatusCode int
	Message    string
}

// ErrorHandler sends an error response used to display error templates.
func ErrorHandler(w http.ResponseWriter, message string, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	errorResponse := ErrorResponse{
		StatusCode: statusCode,
		Message:    message,
	}

	if err := json.NewEncoder(w).Encode(errorResponse); err != nil {
		logger.Error("Failed encoding JSON: %v", err)
		http.Error(w, "An Unexpected Error Occurred. Try Again Later", http.StatusInternalServerError)
		return
	}
}
