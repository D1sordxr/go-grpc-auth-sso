package engine

import (
	"github.com/gin-gonic/gin"
	"src/internal/config/config"
)

type Engine struct {
	Gin *gin.Engine
}

func NewEngine(config *config.Config) *Engine {
	if config.AppConfig.Mode == "prod" {
		gin.SetMode(gin.ReleaseMode)
	}
	return &Engine{gin.Default()}
}
