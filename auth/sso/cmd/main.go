package main

import (
	loadConfig "github.com/D1sordxr/go-grpc-auth-sso/auth/sso/internal/infrastructure/config"
	loadLogger "github.com/D1sordxr/go-grpc-auth-sso/auth/sso/internal/infrastructure/logger"
	"github.com/D1sordxr/go-grpc-auth-sso/auth/sso/internal/presentation/app"
	"github.com/D1sordxr/go-grpc-auth-sso/auth/sso/internal/presentation/grpc/auth"
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

	gRPCServer := auth.NewGrpcServer()

	app := app.NewApp(logger, gRPCServer, "")
	// TODO: run gRPC-server
}
