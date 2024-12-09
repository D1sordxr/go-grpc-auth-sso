package main

import (
	loadUserCommandsService "github.com/D1sordxr/go-grpc-auth-sso/auth/sso/internal/application"
	loadHandlers "github.com/D1sordxr/go-grpc-auth-sso/auth/sso/internal/application/handlers"
	loadConfig "github.com/D1sordxr/go-grpc-auth-sso/auth/sso/internal/infrastructure/config"
	loadDatabase "github.com/D1sordxr/go-grpc-auth-sso/auth/sso/internal/infrastructure/db"
	loadLogger "github.com/D1sordxr/go-grpc-auth-sso/auth/sso/internal/infrastructure/logger"
	loadTokenService "github.com/D1sordxr/go-grpc-auth-sso/auth/sso/internal/infrastructure/token"
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
	uowManager := loadDatabase.NewUoWManager(database)
	userDAO := loadDatabase.NewUserDAO(database)

	tokenService := loadTokenService.NewTokenService(&cfg.TokenConfig)

	// command service v2
	registerUserHandler := loadHandlers.NewRegisterUserHandler(userDAO, uowManager)
	loginUserHandler := loadHandlers.NewLoginUserHandler(userDAO, uowManager, tokenService)
	isAdminUserHandler := loadHandlers.NewIsAdminUserHandler(userDAO, uowManager)

	userCommandsService := loadUserCommandsService.NewUserCommands(registerUserHandler, loginUserHandler, isAdminUserHandler)
	// command service v2

	// userCommands := loadUserCommands.NewUserCommands(userDAO, uowManager)

	authService := loadGRPCServer.NewUserAuthService(userCommandsService)
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
