package db

import (
	"src/internal/db/models"
)

func (s *Storage) CreateOrder(order models.Order) (models.Order, error) {
	id, err := s.OrderDAO.CreateOrder(order)
	if err != nil {
		return models.Order{}, err
	}
	return s.OrderDAO.GetOrder(id)
}

func (s *Storage) GetOrder(id int) (models.Order, error) {
	return s.OrderDAO.GetOrder(id)
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
