package statusOK

import (
	"github.com/D1sordxr/aviasales/src/internal/presentation/api/v1/controllers/handlers/statusOK"
	"github.com/gin-gonic/gin"
)

type Routes struct {
	RouterGroup *gin.RouterGroup
	Handler     *statusOK.Handler
}

func NewOkRoutes(rg *gin.RouterGroup, h *statusOK.Handler) {
	routes := Routes{
		RouterGroup: rg,
		Handler:     h,
	}
	routes.setupOkRoutes()
}

func (r *Routes) setupOkRoutes() {
	r.RouterGroup.GET("/status", r.Handler.GetStatusOK)
}
