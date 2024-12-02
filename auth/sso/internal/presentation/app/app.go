package app

import (
	"github.com/D1sordxr/go-grpc-auth-sso/auth/sso/internal/infrastructure/config/config"
	"github.com/D1sordxr/go-grpc-auth-sso/auth/sso/internal/presentation/grpc/auth"
	"log"
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

func (a *App) Run() {
	port := a.Config.GRPCConfig.Port
	if err := a.GRPCServer.Run(port); err != nil {
		log.Fatalf("Failed to start gRPC server")
	}
}
