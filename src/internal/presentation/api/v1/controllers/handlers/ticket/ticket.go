package ticket

import (
	"github.com/D1sordxr/aviasales/src/internal/application/ticket/dto"
	"github.com/D1sordxr/aviasales/src/internal/db/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type Handler struct {
	UseCase
}

type ResponseData struct {
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

type UseCase interface {
	GetTickets() ([]models.Ticket, error)
	GetTicketByID(id string) (models.Ticket, error)
	GetTicketByIDDTO(id string) (dto.Ticket, error)
	CreateTicketDTO(t dto.Ticket) (dto.Ticket, error)
	CreateTicket(t models.Ticket) error
	UpdateTicket(t models.Ticket) error
	DeleteTicket(id string) error
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

func (h *Handler) GetTicketByID(c *gin.Context) {
	id := c.Param("id")

	ticket, err := h.UseCase.GetTicketByIDDTO(id)
	if err != nil {
		c.JSON(400, ResponseData{"error", err})
		return
	}

	c.JSON(200, ResponseData{Data: ticket})
}

func (h *Handler) CreateTicket(c *gin.Context) {
	var ticket dto.Ticket

	err := c.BindJSON(&ticket)
	if err != nil {
		c.JSON(400, ResponseData{"error", err})
		return
	}

	ticket, err = h.UseCase.CreateTicketDTO(ticket)
	if err != nil {
		c.JSON(400, ResponseData{"error", err})
		return
	}

	c.JSON(200, ResponseData{"Successfully created!", ticket})
}

// UpdateTicket TODO: fix response data (getting by id after updating)
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

// DeleteTicket TODO: showing deleted data (getting by id then deleting)
func (h *Handler) DeleteTicket(c *gin.Context) {
	id := c.Param("id")

	err := h.UseCase.DeleteTicket(id)
	if err != nil {
		c.JSON(400, ResponseData{"error", err})
		return
	}

	c.JSON(200, ResponseData{Message: "Successfully deleted!"})
}
