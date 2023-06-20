package main

import (
	"backgo/config"
	"backgo/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")
	err := config.Init()
	if err != nil {
		// fmt.Println(err)
		logger.Errf("config initialization error: %v", err)
		return
	}
	router.Initialize()

}
