package engine

import (
	"github.com/D1sordxr/aviasales/src/internal/config/config"
	"github.com/gin-gonic/gin"
)

type Engine struct {
	*gin.Engine
}

type GroupRoutes struct {
	*gin.RouterGroup
}

func NewEngine(config *config.Config) *Engine {
	if config.AppConfig.Mode == "prod" {
		gin.SetMode(gin.ReleaseMode)
	}
	return &Engine{gin.Default()}
}
