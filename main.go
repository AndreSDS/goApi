package main

import (
	"github.com/AndreSDS/goApi/router"
	"github.com/AndreSDS/goApi/config"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")
	// initialize configs
	err := config.Init()

	if err != nil {
		logger.Errorf("Error initializing configs: %v", err)
		return
	}

	router.Initialize()
}