package statusOk

import (
	"aviasales/src/internal/db"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Handler struct {
	DBConn *db.Storage
}

func NewOkHandler(storage *db.Storage) *Handler {
	return &Handler{
		DBConn: storage,
	}
}

func (h *Handler) GetStatusOK(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "OK"})
}
