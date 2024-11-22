package dao

import (
	"github.com/D1sordxr/aviasales/src/internal/application/ticket/dto"
	"github.com/D1sordxr/aviasales/src/internal/db/models"
)

type TicketDAO interface {
	GetTickets() ([]models.Ticket, error)
	GetTicketByID(id string) (models.Ticket, error)
	CreateTicketDTO(t dto.Ticket) (dto.Ticket, error)
	CreateTicket(t models.Ticket) error
	UpdateTicket(t models.Ticket) error
	DeleteTicket(id string) error
}
