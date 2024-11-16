package statusOk

import (
	"aviasales/src/internal/db"
	"github.com/gin-gonic/gin"
	"net/http"
)

type OkHandler struct {
	DBConn *db.Storage
}

func NewOkHandler(storage *db.Storage) *OkHandler {
	return &OkHandler{
		DBConn: storage,
	}
}

func (oh *OkHandler) GetStatusOK(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "OK"})
}
