package main

import (
	"fmt"
	"github.com/monferon/apichka/config"
	"github.com/monferon/apichka/internal/app"
)

func main() {
	cfg, err := config.NewConfig()

	if err != nil {
		fmt.Printf("Could not read config file %v", err)
	}

	app.Run(cfg)
}
