package api

import (
	"aviasales/src/internal/db/models"
	"github.com/gin-gonic/gin"
)

func (s *Server) GetTicket(c *gin.Context) {
	data, err := s.DBConn.GetTicket()
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
	// TODO: in db and api
	c.JSON(418, "Currently in development")
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
