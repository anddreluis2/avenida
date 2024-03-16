package main

import (
	"github.com/anddreluis2/avenida/config"
	"github.com/anddreluis2/avenida/router"
)

var (
	logger config.Logger
)

func main() {
	logger = *config.GetLogger("main")
	//Initialize configs

	err := config.Init()
	if err != nil {
		logger.Errorf("Error initializing configs: %v", err)
		return
	}

	router.Initialize()
}
