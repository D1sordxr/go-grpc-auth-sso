package statusOK

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type OkMessage struct {
	Status string `json:"status"`
}

type Handler struct{}

func NewOkHandler() *Handler {
	return &Handler{}
}

func (h *Handler) GetStatusOK(c *gin.Context) {
	c.JSON(http.StatusOK, OkMessage{"OK"})
}
