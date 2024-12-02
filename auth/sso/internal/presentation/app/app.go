package app

import (
	"github.com/D1sordxr/go-grpc-auth-sso/auth/sso/internal/infrastructure/config/config"
	"google.golang.org/grpc"
	"log"
)

type App struct {
	Config     *config.Config
	GRPCServer *grpc.Server
}

func NewApp(config *config.Config,
	gRPC *gin.Engine) *App {
	return &App{
		Config:   config,
		Storage:  storage,
		Router:   router,
		UseCases: useCases,
	}
}

func (a *App) Run() {
	a.registerRoutes()
	port := ":" + a.Config.API.Port
	if err := a.Router.Run(port); err != nil {
		log.Fatalf("error starting server: %v", err)
	}
}

func (a *App) registerRoutes() {
	// Main path
	api := a.Router.Group("/api")

	// V1 path
	routesV1.NewRoutesV1(api, a.UseCases)
}
