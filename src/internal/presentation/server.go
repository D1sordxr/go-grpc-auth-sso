package presentation

import (
	useCase "github.com/D1sordxr/aviasales/src/internal/application"
	"github.com/D1sordxr/aviasales/src/internal/config/config"
	"github.com/D1sordxr/aviasales/src/internal/db"
	orderHandler "github.com/D1sordxr/aviasales/src/internal/presentation/api/controllers/handlers/order"
	okHandler "github.com/D1sordxr/aviasales/src/internal/presentation/api/controllers/handlers/statusOk"
	ticketHandler "github.com/D1sordxr/aviasales/src/internal/presentation/api/controllers/handlers/ticket"
	orderRoutes "github.com/D1sordxr/aviasales/src/internal/presentation/api/controllers/routes/order"
	okRoutes "github.com/D1sordxr/aviasales/src/internal/presentation/api/controllers/routes/statusOk"
	ticketRoutes "github.com/D1sordxr/aviasales/src/internal/presentation/api/controllers/routes/ticket"
	routesV1 "github.com/D1sordxr/aviasales/src/internal/presentation/api/v1"
	"github.com/gin-gonic/gin"
	"log/slog"
)

type Server struct {
	Config  *config.Config
	Logger  *slog.Logger
	DBConn  *db.Storage
	Router  *gin.Engine
	UseCase *useCase.UseCase
}

func NewServer(storage *db.Storage, router *gin.Engine, cfg *config.Config, logger *slog.Logger, useCase *useCase.UseCase) *Server {
	return &Server{
		Config:  cfg,
		Logger:  logger,
		DBConn:  storage,
		Router:  router,
		UseCase: useCase,
	}
}

func (s *Server) Run() error {
	s.registerRoutes()
	port := ":" + s.Config.APIConfig.Port
	if err := s.Router.Run(port); err != nil {
		return err
	}
	return nil
}

func (s *Server) registerRoutes() {
	// Main path
	api := s.Router.Group("/api")

	// V1 path
	routesV1.NewRoutesV1(api, s.UseCase)

	// Status path
	okHandlers := okHandler.NewOkHandler(s.DBConn)
	okRoutes.NewOkRoutes(api, okHandlers)

	// Orders path
	orderHandlers := orderHandler.NewOrderHandler(s.DBConn)
	orderRoutes.NewOrderRoutes(api, orderHandlers)

	// Tickets path
	ticketHandlers := ticketHandler.NewTicketHandler(s.DBConn)
	ticketRoutes.NewTicketRoutes(api, ticketHandlers)
}
