package order

import (
	"aviasales/src/internal/http/api/controllers/handlers/order"
	"github.com/gin-gonic/gin"
)

type Routes struct {
	RouterGroup *gin.RouterGroup
	Handler     *order.Handler
}

func NewOrderRoutes(rg *gin.RouterGroup, h *order.Handler) {
	routes := &Routes{
		RouterGroup: rg,
		Handler:     h,
	}
	routes.setupOrderRoutes()
}

func (r *Routes) setupOrderRoutes() {
	api := r.RouterGroup.Group("/orders")
	{
		api.POST("/order", r.Handler.CreateOrder)
		api.GET("/order/:id", r.Handler.GetOrder)
	}
}
