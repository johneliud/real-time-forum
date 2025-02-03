package controller

import (
	"fmt"
	"log"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		log.Println("Not found")
		ErrorHandler(w, "Not Found", http.StatusNotFound)
		return
	}

	fmt.Fprint(w, "Welcome to home page.")
}
