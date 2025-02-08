package routes

import (
	"net/http"

	"github.com/johneliud/forum/backend/controller"
	"github.com/johneliud/forum/backend/middleware"
)

func Routes() {
	fs := http.FileServer(http.Dir("./frontend"))
	http.Handle("/frontend/", http.StripPrefix("/frontend/", fs))

	http.HandleFunc("/sign-up", controller.SignupHandler)
	http.HandleFunc("/sign-in", controller.SigninHandler)
	http.HandleFunc("/logout", controller.LogoutHandler)

	http.HandleFunc("/validate", controller.ValidateInputHandler)

	http.Handle("/", middleware.AuthenticateMiddleware(http.HandlerFunc(controller.HomeHandler)))
}
