package main

import (
	loadUserCommands "github.com/D1sordxr/go-grpc-auth-sso/auth/sso/internal/application/commands"
	loadConfig "github.com/D1sordxr/go-grpc-auth-sso/auth/sso/internal/infrastructure/config"
	loadDatabase "github.com/D1sordxr/go-grpc-auth-sso/auth/sso/internal/infrastructure/db"
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

	database := loadDatabase.NewConnection(&cfg.DBConfig)
	userDAO := loadDatabase.NewUserDAO(database)

	userCommands := loadUserCommands.NewUserCommands(userDAO)
	authService := loadGRPCServer.NewUserAuthService(userCommands)
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
