package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/johneliud/forum/backend/routes"
	"github.com/johneliud/forum/backend/util"
	"github.com/johneliud/forum/database"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Printf("Error loading .env: %v\n", err)
		return
	}

	util.LoadCredentials()

	database.InitDB()

	routes.Routes()

	fmt.Println("Server started at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
