package ticket

import (
	"github.com/gin-gonic/gin"
	"src/internal/db"
	"src/internal/db/models"
	"strconv"
)

type Handler struct {
	DBConn *db.Storage
}

func NewTicketHandler(storage *db.Storage) *Handler {
	return &Handler{DBConn: storage}
}

func (h *Handler) GetTickets(c *gin.Context) {
	data, err := h.DBConn.GetTickets()
	if err != nil {
		c.JSON(500, "Error reading ticket: "+err.Error())
		return
	}
	c.JSON(200, data)
}

func (h *Handler) CreateTicket(c *gin.Context) {
	var ticket models.Ticket

	if err := c.BindJSON(&ticket); err != nil {
		c.JSON(500, "Error parsing json: "+err.Error())
		return
	}

	err := h.DBConn.CreateTicket(ticket)
	if err != nil {
		c.JSON(500, "Error creating ticket: "+err.Error())
		return
	}
	c.JSON(200, "Successfully created!")
}

func (h *Handler) UpdateTicket(c *gin.Context) {
	var ticket models.Ticket
	id := c.Param("id")

	if err := c.BindJSON(&ticket); err != nil {
		c.JSON(400, "Error parsing json: "+err.Error())
		return
	}

	pID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.JSON(400, "Invalid ticket ID: "+err.Error())
		return
	}
	ticket.ID = &pID

	err = h.DBConn.UpdateTicket(ticket)
	if err != nil {
		c.JSON(400, "Error updating ticket: "+err.Error())
		return
	}
	c.JSON(200, gin.H{
		"message":      "Successfully updated!",
		"updated_data": ticket})
}

func (h *Handler) DeleteTicket(c *gin.Context) {
	id := c.Param("id")

	err := h.DBConn.DeleteTicket(id)
	if err != nil {
		c.JSON(500, "Error deleting ticket: "+err.Error())
		return
	}
	c.JSON(200, "Successfully deleted!")
}
