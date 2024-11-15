package db

import (
	"aviasales/src/internal/db/models"
)

func (s *Storage) CreateOrder(order models.Order) error {
	return s.OrderDAO.CreateOrder(order)
}

func (s *Storage) GetTickets() ([]models.Ticket, error) {
	return s.TicketDAO.GetTickets()
}

func (s *Storage) CreateTicket(t models.Ticket) error {
	return s.TicketDAO.CreateTicket(t)
}

func (s *Storage) UpdateTicket(t models.Ticket) error {
	return s.TicketDAO.UpdateTicket(t)

}

func (s *Storage) DeleteTicket(id string) error {
	return s.TicketDAO.DeleteTicket(id)
}
