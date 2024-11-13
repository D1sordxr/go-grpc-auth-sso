package main

import (
	loadConfig "aviasales/src/internal/config"
	"aviasales/src/internal/db"
	"aviasales/src/internal/http/api"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {

	cfg, err := loadConfig.NewConfig()
	if err != nil {
		log.Fatalf("Failed init config: %v", err)
	}

	storage, err := db.NewDB(cfg.DBConfig.DSN)
	if err != nil {
		log.Fatalf("Failed to connect DB: %v", err)
	}

	_ = storage

	router := gin.Default()
	server := api.NewServer(storage, router)
	api.Setup(server)

	if err = server.Run(); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}

}
