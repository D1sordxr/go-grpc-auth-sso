package api

import (
	"aviasales/src/internal/config"
	"aviasales/src/internal/db"
	ok "aviasales/src/internal/http/api/controllers/handlers/statusOk"
	"aviasales/src/internal/http/api/controllers/routes/statusOk"
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
	api := s.Router.Group("/api")
	okHandler := ok.NewOkHandler(s.DBConn)
	statusOk.NewOkRoutes(api, okHandler)
}
