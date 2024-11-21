package ticket

import (
	"github.com/D1sordxr/aviasales/src/internal/application/ticket/dto"
	"github.com/D1sordxr/aviasales/src/internal/db/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Handler struct {
	UseCase UseCase
}

type ResponseData struct {
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

type UseCase interface {
	GetTickets() ([]models.Ticket, error)
	CreateTicket(t models.Ticket) error
	//UpdateTicket(t models.Ticket) error
	//DeleteTicket(id string) error
}

func NewTicketHandler(useCase UseCase) *Handler {
	return &Handler{UseCase: useCase}

}

func (h *Handler) GetTickets(c *gin.Context) {
	data, err := h.UseCase.GetTickets()
	if err != nil {
		c.JSON(400, ResponseData{"error", err})
		return
	}

	c.JSON(http.StatusOK, ResponseData{Data: data})
}

func (h *Handler) CreateTicket(c *gin.Context) {
	var ticket dto.Ticket

	err := c.BindJSON(&ticket)
	if err != nil {
		c.JSON(400, ResponseData{"error", err})
		return
	}

	mTicket := ticket.ToModel()
	err = h.UseCase.CreateTicket(mTicket)
	if err != nil {
		c.JSON(400, ResponseData{"error", err})
		return
	}

	c.JSON(200, ResponseData{"Successfully created!", ticket})
}
