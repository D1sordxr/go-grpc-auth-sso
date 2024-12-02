package app

import (
	"google.golang.org/grpc"
	"log/slog"
)

type App struct {
	log        *slog.Logger
	gRPCServer *grpc.Server
	port       string
}

func NewApp(log *slog.Logger, grpc *grpc.Server, port string) *App {
	return &App{
		log:        log,
		gRPCServer: grpc,
		port:       port,
	}
}
