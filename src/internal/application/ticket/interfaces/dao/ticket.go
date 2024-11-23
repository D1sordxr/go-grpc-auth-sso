package dao

import (
	"github.com/D1sordxr/aviasales/src/internal/application/ticket/dto"
	"github.com/D1sordxr/aviasales/src/internal/db/models"
)

type TicketDAO interface {
	GetTickets() ([]models.Ticket, error)
	GetTicketsDTO() (dto.Tickets, error)
	GetTicketByID(id string) (models.Ticket, error)
	GetTicketByIDDTO(id string) (dto.Ticket, error)
	CreateTicket(t models.Ticket) error
	CreateTicketDTO(t dto.Ticket) (dto.Ticket, error)
	UpdateTicket(t models.Ticket) error
	UpdateTicketDTO(ticket dto.Ticket) (dto.Ticket, error)
	DeleteTicket(id string) error
	DeleteTicketDTO(id string) (dto.Ticket, error)
}
