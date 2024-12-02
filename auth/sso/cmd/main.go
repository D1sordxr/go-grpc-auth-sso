package main

import (
	loadConfig "github.com/D1sordxr/go-grpc-auth-sso/auth/sso/internal/infrastructure/config"
	loadLogger "github.com/D1sordxr/go-grpc-auth-sso/auth/sso/internal/infrastructure/logger"
	"log"
	"log/slog"
)

func main() {
	cfg, err := loadConfig.NewConfig()
	if err != nil {
		log.Fatalf("failed init config: %v", err)
	}

	logger := loadLogger.NewLogger(cfg)
	logger.Info("starting application", slog.String("mode", cfg.AppConfig.Mode))

	//server := loadGRPC.NewGrpcServer()

	// TODO: init app
	
	// TODO: run gRPC-server
}
