package api

//
//import (
//	"aviasales/src/internal/db/models"
//	"github.com/gin-gonic/gin"
//	"net/http"
//	"strconv"
//)
//
//func (s *Server) CreateOrder(c *gin.Context) {
//	var order models.Order
//	err := c.BindJSON(&order)
//	if err != nil {
//		c.JSON(400, "Error parsing json: "+err.Error())
//		return
//	}
//
//	data, err := s.DBConn.CreateOrder(order)
//	if err != nil {
//		c.JSON(400, "Error creating order: "+err.Error())
//		return
//	}
//
//	c.JSON(http.StatusCreated, gin.H{
//		"message":    "Successfully created!",
//		"order_data": data,
//	})
//}
//
//func (s *Server) GetOrder(c *gin.Context) {
//	strID := c.Param("id")
//	id, err := strconv.ParseInt(strID, 0, 64)
//	if err != nil {
//		c.JSON(400, "can't parse id")
//		return
//	}
//	data, err := s.DBConn.GetOrder(int(id))
//
//	c.JSON(200, gin.H{"order_data": data})
//}
