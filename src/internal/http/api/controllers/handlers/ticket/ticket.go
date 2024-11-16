package ticket

import (
	"aviasales/src/internal/db"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	DBConn *db.Storage
}

func NewTicketHandler(storage *db.Storage) *Handler {
	return &Handler{DBConn: storage}
}

func (h *Handler) CreateTicket(c *gin.Context) {
	c.JSON(418, "api tickets create is teapot")
}
func (h *Handler) GetTickets(c *gin.Context) {
	c.JSON(418, "api tickets get is teapot")
}
func (h *Handler) UpdateTicket(c *gin.Context) {
	c.JSON(418, "api tickets update is teapot")
}
func (h *Handler) DeleteTicket(c *gin.Context) {
	c.JSON(418, "api tickets delete is teapot")
}
