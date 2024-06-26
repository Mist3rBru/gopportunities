package main

import (
	"gopportunities/config"
	"gopportunities/internal/handler"
	"gopportunities/internal/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.NewLogger("main")

	err := config.Init()
	if err != nil {
		logger.Error("config init error: ", err)
		return
	}

	logger.Info("Server running...")
	handler.Init()
	router.Init()
}
