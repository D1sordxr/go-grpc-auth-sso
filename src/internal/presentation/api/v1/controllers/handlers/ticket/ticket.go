package ticket

import (
	"github.com/D1sordxr/aviasales/src/internal/application/ticket/dto"
	"github.com/gin-gonic/gin"
	"strconv"
)

// TODO: move creating and updating logic to domain

type Handler struct {
	UseCase
}

type ResponseData struct {
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

type UseCase interface {
	GetTicketsDTO() (dto.Tickets, error)
	GetTicketByIDDTO(id string) (dto.Ticket, error)
	CreateTicketDTO(t dto.Ticket) (dto.Ticket, error)
	UpdateTicketDTO(ticket dto.Ticket) (dto.Ticket, error)
	DeleteTicketDTO(id string) (dto.Ticket, error)
}

func NewTicketHandler(useCase UseCase) *Handler {
	return &Handler{UseCase: useCase}

}

func (h *Handler) GetTickets(c *gin.Context) {
	tickets, err := h.UseCase.GetTicketsDTO()
	if err != nil {
		c.JSON(400, ResponseData{Message: "error", Data: err.Error()})
		return
	}

	c.JSON(200, ResponseData{Data: tickets.Tickets})
}

func (h *Handler) GetTicketByID(c *gin.Context) {
	id := c.Param("id")

	ticket, err := h.UseCase.GetTicketByIDDTO(id)
	if err != nil {
		c.JSON(400, ResponseData{Message: "error", Data: err.Error()})
		return
	}

	c.JSON(200, ResponseData{Data: ticket})
}

func (h *Handler) CreateTicket(c *gin.Context) {
	var ticket dto.Ticket

	err := c.BindJSON(&ticket)
	if err != nil {
		c.JSON(400, ResponseData{Message: "error", Data: err.Error()})
		return
	}

	ticket, err = h.UseCase.CreateTicketDTO(ticket)
	if err != nil {
		c.JSON(400, ResponseData{Message: "error", Data: err.Error()})
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
		c.JSON(400, ResponseData{Message: "error", Data: err.Error()})
		return
	}

	parsedID, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		c.JSON(400, ResponseData{Message: "error", Data: err.Error()})
		return
	}
	ticket.ID = new(int)
	*ticket.ID = int(parsedID)

	ticket, err = h.UseCase.UpdateTicketDTO(ticket)
	if err != nil {
		c.JSON(400, ResponseData{Message: "error", Data: err.Error()})
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
		c.JSON(400, ResponseData{Message: "error", Data: err.Error()})
		return
	}

	c.JSON(200, ResponseData{
		Message: "Successfully deleted!",
		Data:    deletedTicket,
	})
}
