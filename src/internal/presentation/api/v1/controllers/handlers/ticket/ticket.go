package ticket

import (
	"github.com/D1sordxr/aviasales/src/internal/application/ticket/dto"
	"github.com/D1sordxr/aviasales/src/internal/db/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
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
	UpdateTicket(t models.Ticket) error
	//DeleteTicket(id string) error
}

func NewTicketHandler(useCase UseCase) *Handler {
	return &Handler{UseCase: useCase}

}

func (h *Handler) GetTickets(c *gin.Context) {
	tickets, err := h.UseCase.GetTickets()
	if err != nil {
		c.JSON(400, ResponseData{"error", err})
		return
	}

	c.JSON(http.StatusOK, ResponseData{Data: tickets})
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

func (h *Handler) UpdateTicket(c *gin.Context) {
	var ticket dto.Ticket
	id := c.Param("id")

	err := c.BindJSON(&ticket)
	if err != nil {
		c.JSON(400, ResponseData{"error", err})
		return
	}

	pID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.JSON(400, ResponseData{"error", err})
		return
	}

	mTicket := ticket.ToModel()
	mTicket.ID = &pID

	err = h.UseCase.UpdateTicket(mTicket)
	if err != nil {
		c.JSON(400, ResponseData{"error", err})
		return
	}

	c.JSON(200, ResponseData{"Successfully updated!", ticket})
}
