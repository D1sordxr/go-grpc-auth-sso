package order

import (
	"aviasales/src/internal/http/api"
	"github.com/gin-gonic/gin"
)

type HandlerOrder struct {
	Server *api.Server
	ctx    *gin.Context
}

//
//func (ho *HandlerOrder) CreateOrder(s *api.Server, c *gin.Context) {
//	var order models.Order
//	err := c.BindJSON(&order)
//	if err != nil {
//		c.JSON(400, "Error parsing json: "+err.Error())
//		return
//	}
//
//	err = s.DBConn.CreateOrder(order)
//	if err != nil {
//		c.JSON(400, "Error creating order: "+err.Error())
//		return
//	}
//
//	c.JSON(http.StatusCreated, gin.H{
//		"message":    "Successfully created!",
//		"order_data": order,
//	})
//}
