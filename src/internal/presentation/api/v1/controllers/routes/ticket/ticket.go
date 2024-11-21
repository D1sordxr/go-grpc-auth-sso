package ticket

import (
	"github.com/D1sordxr/aviasales/src/internal/presentation/api/v1/controllers/handlers/ticket"
	"github.com/gin-gonic/gin"
)

type Routes struct {
	RouterGroup *gin.RouterGroup
	Handler     *ticket.Handler
}

func NewTicketRoutes(rg *gin.RouterGroup, h *ticket.Handler) {
	routes := Routes{
		RouterGroup: rg,
		Handler:     h,
	}
	routes.setupTicketRoutes()
}

func (r *Routes) setupTicketRoutes() {
	tickets := r.RouterGroup.Group("/tickets")
	{
		tickets.GET("/ticket", r.Handler.GetTickets)
		tickets.POST("/ticket", r.Handler.CreateTicket)
		tickets.PATCH("/ticket/:id", r.Handler.UpdateTicket)
	}
}
