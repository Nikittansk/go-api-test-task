package main

import (
	"log"

	"github.com/Nikittansk/go-api-test-task/config"
	"github.com/Nikittansk/go-api-test-task/database"
	"github.com/Nikittansk/go-api-test-task/routes"
)

func main() {
	// Reading config file
	cfg, err := config.ReadConfig("config.json")
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	// Initialize Database
	database.Connect(cfg)
	// Initialize Router
	router := routes.InitRouter(cfg)
	router.Run(cfg.Port)
}
