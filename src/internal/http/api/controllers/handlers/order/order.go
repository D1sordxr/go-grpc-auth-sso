package order

import (
	"aviasales/src/internal/db"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	DBConn *db.Storage
}

func NewOrderHandler(storage *db.Storage) *Handler {
	return &Handler{DBConn: storage}
}

func (h *Handler) CreateOrder(c *gin.Context) {
	c.JSON(418, "in progress")
}

func (h *Handler) GetOrder(c *gin.Context) {
	c.JSON(418, "teapot's progress")
}
