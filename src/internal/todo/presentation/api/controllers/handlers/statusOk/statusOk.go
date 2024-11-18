package statusOk

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"src/internal/db"
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
