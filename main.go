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
	var err error

	err = logger.NewLogger("data/app.log", slog.LevelDebug)
	if err != nil {
		panic(err)
	}

	err = util.LoadEnv(".env")
	if err != nil {
		logger.Error("Failed loading .env file: %v", err)
		return
	}

	database.InitDB()
	routes.Routes()

	port, err := util.ValidatePort()
	if err != nil {
		logger.Error("Error validating port: %v", err)
	}

	serverMessage := fmt.Sprintf("Server started at http://localhost%v", port)
	logger.Info(serverMessage)
	http.ListenAndServe(port, nil)
}
