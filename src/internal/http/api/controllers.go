package api

import (
	"aviasales/src/internal/db/models"
	"aviasales/src/internal/http/api/order"
	"github.com/gin-gonic/gin"
	"strconv"
)

func (s *Server) CreateOrder(c *gin.Context) {
	order.CreateOrder(s, c)
}

func (s *Server) GetTickets(c *gin.Context) {
	data, err := s.DBConn.GetTickets()
	if err != nil {
		c.JSON(500, "Error reading ticket: "+err.Error())
		return
	}
	c.JSON(200, data)
}

func (s *Server) CreateTicket(c *gin.Context) {
	var ticket models.Ticket

	if err := c.BindJSON(&ticket); err != nil {
		c.JSON(500, "Error parsing json: "+err.Error())
		return
	}

	err := s.DBConn.CreateTicket(ticket)
	if err != nil {
		c.JSON(500, "Error creating ticket: "+err.Error())
		return
	}
	c.JSON(200, "Successfully created!")
}

func (s *Server) UpdateTicket(c *gin.Context) {
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

	err = s.DBConn.UpdateTicket(ticket)
	if err != nil {
		c.JSON(400, "Error updating ticket: "+err.Error())
		return
	}
	c.JSON(200, gin.H{
		"message":      "Successfully updated!",
		"updated_data": ticket})
}

func (s *Server) DeleteTicket(c *gin.Context) {
	id := c.Param("id")

	err := s.DBConn.DeleteTicket(id)
	if err != nil {
		c.JSON(500, "Error creating ticket: "+err.Error())
		return
	}
	c.JSON(200, "Successfully deleted!")
}
