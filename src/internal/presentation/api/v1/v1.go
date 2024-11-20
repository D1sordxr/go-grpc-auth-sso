package v1

import (
	useCase "github.com/D1sordxr/aviasales/src/internal/application"
	okHandler "github.com/D1sordxr/aviasales/src/internal/presentation/api/v1/controllers/handlers/statusOK"
	okRoutes "github.com/D1sordxr/aviasales/src/internal/presentation/api/v1/controllers/routes/statusOK"
	"github.com/gin-gonic/gin"
)

type RoutesV1 struct {
	RouterGroup *gin.RouterGroup
	UseCase     *useCase.UseCase
}

func NewRoutesV1(rg *gin.RouterGroup, uc *useCase.UseCase) {
	routes := &RoutesV1{
		RouterGroup: rg,
		UseCase:     uc,
	}
	routes.setupRoutesV1()
}

func (r *RoutesV1) setupRoutesV1() {
	// Main path
	v1 := r.RouterGroup.Group("/v1")

	// Status path
	okHandlers := okHandler.NewOkHandler()
	okRoutes.NewOkRoutes(v1, okHandlers)
}
