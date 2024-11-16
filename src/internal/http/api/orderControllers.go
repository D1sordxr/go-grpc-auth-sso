package api

import (
	"aviasales/src/internal/db/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (s *Server) CreateOrder(c *gin.Context) {
	var order models.Order
	err := c.BindJSON(&order)
	if err != nil {
		c.JSON(400, "Error parsing json: "+err.Error())
		return
	}

	data, err := s.DBConn.CreateOrder(order)
	if err != nil {
		c.JSON(400, "Error creating order: "+err.Error())
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message":    "Successfully created!",
		"order_data": data,
	})
}

// TODO: func (s *Server) GetOrder(c *gin.Context) {}
// TODO: func (s *Server) PayOrder(c *gin.Context) {}
// TODO: func (s *Server) DeleteOrder(c *gin.Context) {}
