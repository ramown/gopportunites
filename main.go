package main

import (
	"github.com/ramown/gopportunites/config"
	"github.com/ramown/gopportunites/router"
)

var (
	logger config.Logger
)

func main() {
	logger = *config.GetLogger("main")
	// Initialize configs
	if err := config.Init(); err != nil {
		logger.Errorf("config initialization erro: %v", err)
		return
	}

	router.Initialize()
}
