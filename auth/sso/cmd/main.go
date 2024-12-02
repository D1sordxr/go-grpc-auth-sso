package main

import (
	loadConfig "github.com/D1sordxr/go-grpc-auth-sso/auth/sso/internal/infrastructure/config"
	loadLogger "github.com/D1sordxr/go-grpc-auth-sso/auth/sso/internal/infrastructure/logger"
	loadApp "github.com/D1sordxr/go-grpc-auth-sso/auth/sso/internal/presentation/app"
	loadGRPCServer "github.com/D1sordxr/go-grpc-auth-sso/auth/sso/internal/presentation/grpc/auth"
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

	// TODO: implement services
	gRPCServer := loadGRPCServer.NewGRPCServer()

	app := loadApp.NewApp(cfg, logger.Logger, gRPCServer)
	if err = app.Run(); err != nil {
		logger.Error("Failed to run application", slog.String("error", err.Error()))
	}
}
