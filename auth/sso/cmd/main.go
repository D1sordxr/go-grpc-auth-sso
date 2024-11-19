package main

import (
	"log"
	"log/slog"
	loadConfig "sso/internal/infrastructure/config"
	loadLogger "sso/internal/infrastructure/logger"
)

func main() {
	cfg, err := loadConfig.NewConfig()
	if err != nil {
		log.Fatalf("failed init config: %v", err)
	}

	logger := loadLogger.NewLogger(cfg)
	logger.Info("starting application", slog.String("mode", cfg.AppConfig.Mode))

	// TODO: init app

	// TODO: run gRPC-server
}
