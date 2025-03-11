package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"github.com/johneliud/real-time-forum/backend/logger"
	"github.com/johneliud/real-time-forum/backend/routes"
	"github.com/johneliud/real-time-forum/backend/util"
	"github.com/johneliud/real-time-forum/database"
)

func main() {
	var err error

	_ = os.Mkdir("data", 0o700)

	err = logger.NewLogger("data/app.log", slog.LevelDebug)
	if err != nil {
		panic(err)
	}

	port, err := util.ValidatePort()
	if err != nil {
		logger.Error("Error validating port:", err)
		return
	}

	err = util.LoadEnv(".env")
	if err != nil {
		logger.Error("Failed loading .env file:", err)
		return
	}

	database.InitDB()
	routes.Routes()

	serverMessage := fmt.Sprintf("Server started at http://localhost%v", port)
	logger.Info(serverMessage)
	http.ListenAndServe(port, nil)
}
