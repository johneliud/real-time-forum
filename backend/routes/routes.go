package routes

import (
	"net/http"

	"github.com/johneliud/real-time-forum/backend/controller"
)

func Routes() {
	fs := http.FileServer(http.Dir("./frontend"))
	http.Handle("/frontend/", http.StripPrefix("/frontend/", fs))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./frontend/templates/index.html")
	})

	// API routes
	http.HandleFunc("/api/sign-up", controller.SignupHandler)
	http.HandleFunc("/api/validate", controller.ValidateInputHandler)

	http.HandleFunc("/api/sign-in", controller.SigninHandler)
}
