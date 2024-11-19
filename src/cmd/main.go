package main

import (
	loadConfig "github.com/D1sordxr/aviasales/src/internal/config"
	loadDB "github.com/D1sordxr/aviasales/src/internal/db"
	loadLogger "github.com/D1sordxr/aviasales/src/internal/logger"
	loadServer "github.com/D1sordxr/aviasales/src/internal/presentation"
	loadRouter "github.com/D1sordxr/aviasales/src/internal/presentation/api/engine"
	"log"
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
