package main

import (
	"github.com/bensonopisa/tasktracker/config"
	"github.com/bensonopisa/tasktracker/internal/app"
)

func main() {

	cfg := config.ReadConfig()

	app.Run(cfg)
}
