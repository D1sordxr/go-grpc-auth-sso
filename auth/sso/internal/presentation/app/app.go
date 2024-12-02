package app

import (
	"fmt"
	"github.com/D1sordxr/go-grpc-auth-sso/auth/sso/internal/infrastructure/config/config"
	"github.com/D1sordxr/go-grpc-auth-sso/auth/sso/internal/presentation/grpc/auth"
	"log/slog"
	"net"
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
	const operation = "gRPC.Run"
	port := a.Config.GRPCConfig.Port
	if err := a.GRPCServer.RegisterServer(); err != nil {
		return fmt.Errorf("failed to register gRPC server: %w", err)
	}

	l, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return fmt.Errorf("%s: %w", operation, err)
	}

	a.Logger.Info("grpc server started", slog.String("port", fmt.Sprintf("%d", port)))

	if err = a.GRPCServer.Server.Serve(l); err != nil {
		return fmt.Errorf("%s: %w", operation, err)
	}

	return nil
}
