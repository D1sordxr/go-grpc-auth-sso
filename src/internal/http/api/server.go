package api

import (
	"aviasales/src/internal/db"
	"github.com/gin-gonic/gin"
)

type Server struct {
	DBConn *db.Storage
	Router *gin.Engine
}

func NewServer(storage *db.Storage, router *gin.Engine) *Server {
	//order.NewServerOrder(storage, router)
	return &Server{
		DBConn: storage,
		Router: router,
	}
}

func (s *Server) Run() error {
	if err := s.Router.Run(); err != nil {
		return err
	}
	return nil
}
