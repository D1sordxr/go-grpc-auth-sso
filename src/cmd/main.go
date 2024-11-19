package main

import (
	"log"
	loadConfig "src/internal/config"
	loadDB "src/internal/db"
	loadLogger "src/internal/logger"
	loadServer "src/internal/presentation"
	loadRouter "src/internal/presentation/api/engine"
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

	// TODO: useCase := loadUseCase.NewUseCase(storage.TicketDAO, storage.OrderDAO)

	router := loadRouter.NewEngine(cfg).Engine

	server := loadServer.NewServer(storage, router, cfg, logger)

	if err = server.Run(); err != nil {
		logger.Log().Err(err).Msg("failed to start http server")
	}
}
