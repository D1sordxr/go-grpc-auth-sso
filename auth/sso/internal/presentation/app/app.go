package app

import (
	"fmt"
	"github.com/D1sordxr/go-grpc-auth-sso/auth/sso/internal/infrastructure/config/config"
	"github.com/D1sordxr/go-grpc-auth-sso/auth/sso/internal/presentation/grpc/auth"
	"log/slog"
	"net"
	"os"
	"os/signal"
	"syscall"
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
	var err error
	errorsChannel := make(chan error, 1)

	go func() {
		if err = a.gRPCServerRun(); err != nil {
			errorsChannel <- err
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	select {
	case <-stop:
		a.Logger.Info("Stopping application...", slog.String("signal", "stop"))
	case err = <-errorsChannel:
		a.Logger.Error("Application encountered an error", slog.String("error", err.Error()))
	}

	a.GRPCServer.Down()
	a.Logger.Info("Gracefully stopped")
}

func (a *App) gRPCServerRun() error {
	const operation = "gRPCServer.Run"
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
