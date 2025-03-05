package controller

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/johneliud/real-time-forum/database"
)

/*
ValidateInputHandler checks if a name or email already exists in the database.
*/
func ValidateInputHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/validate" {
		log.Printf("Failed finding %q route.\n", r.URL.Path)
		ErrorHandler(w, "Not Found", http.StatusNotFound)
		return
	}

	if r.Method != http.MethodGet {
		log.Printf("Invalid request method: %v\n", r.Method)
		ErrorHandler(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	if err := r.ParseForm(); err != nil {
		log.Printf("Failed parsing form: %v\n", err)
		ErrorHandler(w, "Bad Request", http.StatusBadRequest)
		return
	}

	name, email := strings.TrimSpace(r.FormValue("name")), strings.TrimSpace(r.FormValue("email"))
	var query, value string

	if name != "" {
		query = "SELECT id FROM users WHERE name = ?"
		value = name
	} else if email != "" {
		query = "SELECT id FROM users WHERE email = ?"
		value = email
	} else {
		log.Println("Invalid input provided.")
		ErrorHandler(w, "Bad Request", http.StatusBadRequest)
		return
	}

	var userID int
	err := database.DB.QueryRow(query, value).Scan(&userID)
	// Provided credentials are unique
	if err == sql.ErrNoRows {
		json.NewEncoder(w).Encode(map[string]bool{"available": true})
		return
	} else if err != nil {
		log.Printf("Failed quering databse: %v\n", err)
		ErrorHandler(w, "Something Unexpected Happened. Try Again Later", http.StatusInternalServerError)
		return
	}
	// Provided credentials are taken
	json.NewEncoder(w).Encode(map[string]bool{"available": false})
}
