package dao

import "src/internal/db/models"

type TicketDAO interface {
	GetTickets() ([]models.Ticket, error)
	CreateTicket(t models.Ticket) error
	UpdateTicket(t models.Ticket) error
	DeleteTicket(id string) error
}
