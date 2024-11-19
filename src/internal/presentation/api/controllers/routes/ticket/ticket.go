package ticket

import (
	"github.com/D1sordxr/aviasales/src/internal/presentation/api/controllers/handlers/ticket"
	"github.com/gin-gonic/gin"
)

type Routes struct {
	RouterGroup *gin.RouterGroup
	Handler     *ticket.Handler
}

func NewTicketRoutes(rg *gin.RouterGroup, h *ticket.Handler) {
	routes := &Routes{
		RouterGroup: rg,
		Handler:     h,
	}
	routes.setupTicketRoutes()
}

func (r *Routes) setupTicketRoutes() {
	api := r.RouterGroup.Group("/tickets")
	{
		api.GET("/ticket", r.Handler.GetTickets)
		api.POST("/ticket", r.Handler.CreateTicket)
		api.PATCH("/ticket/:id", r.Handler.UpdateTicket)
		api.DELETE("/ticket/:id", r.Handler.DeleteTicket)
	}
}
