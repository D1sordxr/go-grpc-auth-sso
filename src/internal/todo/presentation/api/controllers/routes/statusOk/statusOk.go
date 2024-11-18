package statusOk

import (
	"github.com/gin-gonic/gin"
	ok "src/internal/http/api/controllers/handlers/statusOk"
)

type OkRoutes struct {
	RouterGroup *gin.RouterGroup
	Handler     *ok.Handler
}

func NewOkRoutes(rg *gin.RouterGroup, handler *ok.Handler) {
	routes := &OkRoutes{
		RouterGroup: rg,
		Handler:     handler,
	}
	routes.setupOkRoutes()
}

func (r *OkRoutes) setupOkRoutes() {
	r.RouterGroup.GET("/status", r.Handler.GetStatusOK)
}
