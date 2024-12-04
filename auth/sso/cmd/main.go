package main

import (
	loadConfig "github.com/D1sordxr/go-grpc-auth-sso/auth/sso/internal/infrastructure/config"
	loadLogger "github.com/D1sordxr/go-grpc-auth-sso/auth/sso/internal/infrastructure/logger"
	loadApp "github.com/D1sordxr/go-grpc-auth-sso/auth/sso/internal/presentation/app"
	loadGRPCServer "github.com/D1sordxr/go-grpc-auth-sso/auth/sso/internal/presentation/grpc/auth"
	"log"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	cfg, err := loadConfig.NewConfig()
	if err != nil {
		log.Fatalf("failed init config: %v", err)
	}

	logger := loadLogger.NewLogger(cfg)
	logger.Info("starting application", slog.String("mode", cfg.AppConfig.Mode))

	// TODO: database

	// TODO: implement services
	authService := loadGRPCServer.NewUserAuthService()
	gRPCServer := loadGRPCServer.NewGRPCServer(authService)

	app := loadApp.NewApp(cfg, logger.Logger, gRPCServer)

	errorsChannel := make(chan error, 1)
	go func() {
		if err = app.Run(); err != nil {
			errorsChannel <- err
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	select {
	case <-stop:
		logger.Info("Stopping application...", slog.String("signal", "stop"))
	case err = <-errorsChannel:
		logger.Error("Application encountered an error", slog.String("error", err.Error()))
	}

	app.GRPCServer.Down()
	logger.Info("Gracefully stopped")
}
