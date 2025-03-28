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

	_ = os.Mkdir("files", 0o700)

	err = logger.NewLogger("files/app.log", slog.LevelDebug)
	if err != nil {
		panic(err)
	}

	port, err := util.ValidatePort()
	if err != nil {
		logger.Error("Error validating port", "err", err)
		return
	}

	err = util.LoadEnv(".env")
	if err != nil {
		logger.Error("Failed loading .env file", "err", err)
		return
	}

	database.InitDB()
	mux := routes.Routes()

	logger.Info(fmt.Sprintf("Server started at http://localhost%s", port))
	if err := http.ListenAndServe(port, mux); err != nil {
		logger.Error("Server failed to start", "err", err)
		return
	}
}
