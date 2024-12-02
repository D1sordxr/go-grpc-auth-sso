package app

import (
	"google.golang.org/grpc"
	"log/slog"
)

type App struct {
	Logger     *slog.Logger
	GRPCServer *grpc.Server
	Port       string
}

func NewApp(log *slog.Logger, grpc *grpc.Server, port string) *App {
	return &App{
		GRPCServer: grpc.NewServer(),
	}
}
