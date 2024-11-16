package statusOk

import (
	ok "aviasales/src/internal/http/api/controllers/handlers/statusOk"
	"github.com/gin-gonic/gin"
)

type OkRoutes struct {
	RouterGroup *gin.RouterGroup
	Handler     *ok.OkHandler
}

func NewOkRoutes(rg *gin.RouterGroup, handler *ok.OkHandler) {
	routes := &OkRoutes{
		RouterGroup: rg,
		Handler:     handler,
	}
	routes.setupOkRoutes()
}

func (r *OkRoutes) setupOkRoutes() {
	r.RouterGroup.GET("/status", r.Handler.GetStatusOK)
}
