package routes

import (
	"net/http"

	"github.com/johneliud/real-time-forum/backend/controller"
)

func Routes() {
	fs := http.FileServer(http.Dir("./frontend"))
	http.Handle("/frontend/", http.StripPrefix("/frontend/", fs))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./frontend/index.html")
	})

	http.HandleFunc("/sign-up", controller.SignupHandler)

	http.HandleFunc("/validate", controller.ValidateInputHandler)
}
