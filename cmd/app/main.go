package main

import (
	"apichka/config"
	"apichka/internal/app"
	"fmt"
)

func main() {
	cfg, err := config.NewConfig()

	if err != nil {
		fmt.Printf("Could not read config file %v", err)
	}

	app.Run(cfg)
}
