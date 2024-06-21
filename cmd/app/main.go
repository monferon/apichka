package main

import (
	"apichka/config"
	"apichka/internal/app"
	"errors"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		errors.New("Error loading config: " + err.Error())
	}
	app.Run(cfg)
}
