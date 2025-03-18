package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"text/template"
	"github.com/johneliud/real-time-forum/database"
)

/*
HomeHandler serves the index page of the web page.
*/
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		log.Println("Not found")
		ErrorHandler(w, "Not Found", http.StatusNotFound)
		return
	}

	tmpl, err := template.ParseFiles("frontend/templates/index.html")
	if err != nil {
		log.Printf("Failed parsing template: %v\n", err)
		ErrorHandler(w, "Page Not Found", http.StatusNotFound)
		return
	}

	err = tmpl.Execute(w, nil)
	if err != nil {
		log.Printf("Failed to execute template: %v\n", err)
		ErrorHandler(w, "Something Unexpected Happened. Try Again Later", http.StatusInternalServerError)
		return
	}
}

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
