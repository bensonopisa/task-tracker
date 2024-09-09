package app

import (
	"fmt"
	"log"

	"github.com/bensonopisa/tasktracker/config"
	"github.com/bensonopisa/tasktracker/internal/database"
)

func Run(cfg *config.Config) {
	fmt.Println(cfg.AppCfg.DbConfig)
	_, err := database.Connect(cfg.AppCfg.DbConfig)

	if err != nil {
		log.Panicf("Error while connecting to db: %v", err)
	}

	log.Println("Application has been setup correctly this far.")
}
