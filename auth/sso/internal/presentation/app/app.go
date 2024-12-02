package app

import (
	"fmt"
	"github.com/D1sordxr/go-grpc-auth-sso/auth/sso/internal/infrastructure/config/config"
	"github.com/D1sordxr/go-grpc-auth-sso/auth/sso/internal/presentation/grpc/auth"
	"log/slog"
)

type App struct {
	Config     *config.Config
	Logger     *slog.Logger
	GRPCServer *auth.Server
}

func NewApp(config *config.Config,
	logger *slog.Logger,
	gRPC *auth.Server) *App {
	return &App{
		Config:     config,
		Logger:     logger,
		GRPCServer: gRPC,
	}
}

func (a *App) Run() error {
	port := a.Config.GRPCConfig.Port
	if err := a.GRPCServer.Run(port); err != nil {
		return fmt.Errorf("failed to start gRPC server: %w", err)
	}
	return nil
}
