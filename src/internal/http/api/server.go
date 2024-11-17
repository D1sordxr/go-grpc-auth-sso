package api

import (
	"aviasales/src/internal/config"
	"aviasales/src/internal/db"
	orderHandler "aviasales/src/internal/http/api/controllers/handlers/order"
	okHandler "aviasales/src/internal/http/api/controllers/handlers/statusOk"
	ticketHandler "aviasales/src/internal/http/api/controllers/handlers/ticket"
	orderRoutes "aviasales/src/internal/http/api/controllers/routes/order"
	okRoutes "aviasales/src/internal/http/api/controllers/routes/statusOk"
	ticketRoutes "aviasales/src/internal/http/api/controllers/routes/ticket"
	"github.com/gin-gonic/gin"
)

type Server struct {
	DBConn *db.Storage
	Router *gin.Engine
	Config *config.Config
}

func NewServer(storage *db.Storage, router *gin.Engine, cfg *config.Config) *Server {
	return &Server{
		DBConn: storage,
		Router: router,
		Config: cfg,
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
