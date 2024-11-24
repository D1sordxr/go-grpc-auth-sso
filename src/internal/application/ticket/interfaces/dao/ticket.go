package dao

import (
	"github.com/D1sordxr/aviasales/src/internal/application/ticket/dto"
)

type TicketDAO interface {
	GetTicketsDTO() (dto.Tickets, error)
	GetTicketByIDDTO(id string) (dto.Ticket, error)
	CreateTicketDTO(t dto.Ticket) (dto.Ticket, error)
	UpdateTicketDTO(ticket dto.Ticket) (dto.Ticket, error)
	DeleteTicketDTO(id string) (dto.Ticket, error)
}
