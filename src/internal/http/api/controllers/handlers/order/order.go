package order

import (
	"aviasales/src/internal/db"
	"aviasales/src/internal/db/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type Handler struct {
	DBConn *db.Storage
}

func NewOrderHandler(storage *db.Storage) *Handler {
	return &Handler{DBConn: storage}
}

func (h *Handler) CreateOrder(c *gin.Context) {
	var order models.Order
	err := c.BindJSON(&order)
	if err != nil {
		c.JSON(400, "Error parsing json: "+err.Error())
		return
	}

	data, err := h.DBConn.CreateOrder(order)
	if err != nil {
		c.JSON(400, "Error creating order: "+err.Error())
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message":    "Successfully created!",
		"order_data": data,
	})
}

func (h *Handler) GetOrder(c *gin.Context) {
	strID := c.Param("id")
	id, err := strconv.ParseInt(strID, 0, 64)
	if err != nil {
		c.JSON(400, "can't parse id")
		return
	}
	data, err := h.DBConn.GetOrder(int(id))

	c.JSON(200, gin.H{"order_data": data})
}

// TODO: func (h *Handler) PayOrder(c *gin.Context) {}
// TODO: func (h *Handler) DeleteOrder(c *gin.Context) {}
