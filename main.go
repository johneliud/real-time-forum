package main

import (
	"fmt"
	"net/http"

	"github.com/johneliud/forum/backend/controller"
	"github.com/johneliud/forum/database"
)

func main() {
	database.InitDB()
	
	fs := http.FileServer(http.Dir("./frontend"))
	http.Handle("/frontend/", http.StripPrefix("/frontend/", fs))

	http.HandleFunc("/sign-up", controller.SignupHandler)

	fmt.Println("Server started at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
