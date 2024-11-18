package main

import (
	"github.com/gin-gonic/gin"
	"log"
	loadConfig "src/internal/config"
	loadDB "src/internal/db"
	"src/internal/http/api"
	loadLogger "src/internal/logger"
)

func main() {
	cfg, err := loadConfig.NewConfig()
	if err != nil {
		log.Fatalf("failed init config: %v", err)
	}

	logger := loadLogger.NewLogger(cfg.LoggerConfig)
	logger.Info().Msg("starting flights order app on mode: " + cfg.AppConfig.Mode)

	storage, err := loadDB.NewDB(&cfg.DBConfig)
	if err != nil {
		logger.Log().Err(err).Msg("failed to connect DB")
	}

	router := gin.Default()
	server := api.NewServer(storage, router, cfg)

	if err = server.Run(); err != nil {
		logger.Log().Err(err).Msg("failed to start http server")
	}
}
