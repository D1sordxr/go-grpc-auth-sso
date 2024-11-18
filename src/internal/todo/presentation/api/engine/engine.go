package engine

import (
	"github.com/gin-gonic/gin"
)

type Engine struct {
	Engine *gin.Engine
}

type RouterGroup struct {
	*gin.RouterGroup
}

func NewEngine() Engine {
	return Engine{gin.Default()}
}
