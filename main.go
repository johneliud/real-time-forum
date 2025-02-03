package main

import (
	"fmt"
	"net/http"

	"github.com/johneliud/forum/backend/controller"
	"github.com/johneliud/forum/backend/middleware"
	"github.com/johneliud/forum/database"
)

func main() {
	database.InitDB()
	
	fs := http.FileServer(http.Dir("./frontend"))
	http.Handle("/frontend/", http.StripPrefix("/frontend/", fs))

	http.HandleFunc("/sign-up", controller.SignupHandler)
	http.HandleFunc("/sign-in", controller.SigninHandler)
	http.HandleFunc("/logout", controller.LogoutHandler)

	http.Handle("/", middleware.AuthenticateMiddleware(http.HandlerFunc(controller.HomeHandler)))

	fmt.Println("Server started at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
