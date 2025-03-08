package main

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/johneliud/real-time-forum/backend/logger"
	"github.com/johneliud/real-time-forum/backend/routes"
	"github.com/johneliud/real-time-forum/backend/util"
	"github.com/johneliud/real-time-forum/database"
)

func main() {
	err := util.LoadEnv(".env")
	if err != nil {
		fmt.Println("Error loading .env file:", err)
		return
	}

	err = logger.NewLogger("app.log", slog.LevelDebug)
	if err != nil {
		panic(err)
	}

	database.InitDB()

	routes.Routes()

	logger.Info("Server started at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
