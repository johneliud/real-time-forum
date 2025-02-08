package main

import (
	"fmt"
	"net/http"

	"github.com/johneliud/forum/backend/routes"
	"github.com/johneliud/forum/database"
)

func main() {
	database.InitDB()

	routes.Routes()

	fmt.Println("Server started at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
