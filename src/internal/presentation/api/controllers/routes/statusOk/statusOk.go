package statusOk

import (
	ok "github.com/D1sordxr/aviasales/src/internal/presentation/api/controllers/handlers/statusOk"
	"github.com/gin-gonic/gin"
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
