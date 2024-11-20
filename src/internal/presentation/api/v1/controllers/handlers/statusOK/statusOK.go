package statusOK

import (
	useCase "github.com/D1sordxr/aviasales/src/internal/application"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Handler struct {
	*useCase.UseCase
}

func NewOkHandler(useCase *useCase.UseCase) *Handler {
	return &Handler{useCase}
}

func (h *Handler) GetStatusOK(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "OK"})
}
