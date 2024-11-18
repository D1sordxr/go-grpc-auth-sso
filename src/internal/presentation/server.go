package presentation

import (
	"github.com/gin-gonic/gin"
	"src/internal/config/config"
	"src/internal/db"
	"src/internal/logger"
	orderHandler "src/internal/presentation/api/controllers/handlers/order"
	okHandler "src/internal/presentation/api/controllers/handlers/statusOk"
	ticketHandler "src/internal/presentation/api/controllers/handlers/ticket"
	orderRoutes "src/internal/presentation/api/controllers/routes/order"
	okRoutes "src/internal/presentation/api/controllers/routes/statusOk"
	ticketRoutes "src/internal/presentation/api/controllers/routes/ticket"
)

type Server struct {
	Config *config.Config
	Logger *logger.Logger
	DBConn *db.Storage
	Router *gin.Engine
}

func NewServer(storage *db.Storage, router *gin.Engine, cfg *config.Config, logger *logger.Logger) *Server {
	return &Server{
		Config: cfg,
		Logger: logger,
		DBConn: storage,
		Router: router,
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
