package api

import (
	"aviasales/src/internal/config"
	"aviasales/src/internal/db"
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
	port := ":" + s.Config.APIConfig.Port
	if err := s.Router.Run(port); err != nil {
		return err
	}
	return nil
}
