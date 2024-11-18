package main

import (
	"github.com/gin-gonic/gin"
	"log"
	loadConfig "src/internal/config"
	"src/internal/db"
	"src/internal/http/api"
)

func main() {
	cfg, err := loadConfig.NewConfig()
	if err != nil {
		log.Fatalf("Failed init config: %v", err)
	}

	storage, err := db.NewDB(&cfg.DBConfig)
	if err != nil {
		log.Fatalf("Failed to connect DB: %v", err)
	}

	router := gin.Default()
	server := api.NewServer(storage, router, cfg)

	if err = server.Run(); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
