package main

import (
	"fmt"
	"net/http"

	"github.com/johneliud/forum/backend/routes"
	"github.com/johneliud/forum/backend/util"
	"github.com/johneliud/forum/database"
)

func main() {
	err := util.LoadEnv(".env")
	if err != nil {
		fmt.Println("Error loading .env file:", err)
		return
	}

	database.InitDB()

	routes.Routes()

	fmt.Println("Server started at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
