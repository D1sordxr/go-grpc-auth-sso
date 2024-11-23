package ticket

import (
	"github.com/D1sordxr/aviasales/src/internal/application/ticket/dto"
	"github.com/D1sordxr/aviasales/src/internal/db/models"
	"github.com/gin-gonic/gin"
	"strconv"
)

type Handler struct {
	UseCase
}

type ResponseData struct {
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Error   error       `json:"error,omitempty"`
}

type UseCase interface {
	GetTickets() ([]models.Ticket, error)
	GetTicketsDTO() (dto.Tickets, error)
	GetTicketByID(id string) (models.Ticket, error)
	GetTicketByIDDTO(id string) (dto.Ticket, error)
	CreateTicketDTO(t dto.Ticket) (dto.Ticket, error)
	CreateTicket(t models.Ticket) error
	UpdateTicket(t models.Ticket) error
	UpdateTicketDTO(ticket dto.Ticket) (dto.Ticket, error)
	DeleteTicket(id string) error
	DeleteTicketDTO(id string) (dto.Ticket, error)
}

func NewTicketHandler(useCase UseCase) *Handler {
	return &Handler{UseCase: useCase}

}

func (h *Handler) GetTickets(c *gin.Context) {
	tickets, err := h.UseCase.GetTicketsDTO()
	if err != nil {
		c.JSON(400, ResponseData{Message: "error", Error: err})
		return
	}

	c.JSON(200, ResponseData{Data: tickets.Tickets})
}

func (h *Handler) GetTicketByID(c *gin.Context) {
	id := c.Param("id")

	ticket, err := h.UseCase.GetTicketByIDDTO(id)
	if err != nil {
		c.JSON(400, ResponseData{Message: "error", Error: err})
		return
	}

	c.JSON(200, ResponseData{Data: ticket})
}

func (h *Handler) CreateTicket(c *gin.Context) {
	var ticket dto.Ticket

	err := c.BindJSON(&ticket)
	if err != nil {
		c.JSON(400, ResponseData{Message: "error", Error: err})
		return
	}

	ticket, err = h.UseCase.CreateTicketDTO(ticket)
	if err != nil {
		c.JSON(400, ResponseData{Message: "error", Error: err})
		return
	}

	c.JSON(200, ResponseData{
		Message: "Successfully created!",
		Data:    ticket,
	})
}

func (h *Handler) UpdateTicket(c *gin.Context) {
	var ticket dto.Ticket
	id := c.Param("id")

	err := c.BindJSON(&ticket)
	if err != nil {
		c.JSON(400, ResponseData{Message: "error", Error: err})
		return
	}

	parsedID, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		c.JSON(400, ResponseData{Message: "error", Error: err})
		return
	}
	ticket.ID = new(int)
	*ticket.ID = int(parsedID)

	ticket, err = h.UseCase.UpdateTicketDTO(ticket)
	if err != nil {
		c.JSON(400, ResponseData{Message: "error", Error: err})
		return
	}
	c.JSON(200, ResponseData{
		Message: "Successfully updated!",
		Data:    ticket,
	})
}

func (h *Handler) DeleteTicket(c *gin.Context) {
	id := c.Param("id")

	deletedTicket, err := h.UseCase.DeleteTicketDTO(id)
	if err != nil {
		c.JSON(400, ResponseData{Message: "error", Error: err})
		return
	}

	c.JSON(200, ResponseData{
		Message: "Successfully deleted!",
		Data:    deletedTicket,
	})
}
