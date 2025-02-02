package controller

import (
	"log"
	"net/http"
	"text/template"
)

type ErrorPage struct {
	StatusCode int
	Message    string
}

/*
ErrorHandler parses an error template when displaying error information.
*/
func ErrorHandler(w http.ResponseWriter, message string, statusCode int) {
	tmpl, err := template.ParseFiles("frontend/templates/error-page.html")
	if err != nil {
		log.Printf("Failed parsing template: %v\n", err)
		http.Error(w, "Page Not Found", http.StatusNotFound)
		return
	}

	data := ErrorPage{
		StatusCode: statusCode,
		Message:    message,
	}

	if err = tmpl.Execute(w, data); err != nil {
		log.Printf("Failed executing template: %v\n", err)
		http.Error(w, "Something Unexpected Happened. Try Again Later", http.StatusInternalServerError)
		return
	}
}
