package application

import (
	order "github.com/D1sordxr/aviasales/src/internal/application/order/interfaces/dao"
	"github.com/D1sordxr/aviasales/src/internal/application/ticket/dto"
	ticket "github.com/D1sordxr/aviasales/src/internal/application/ticket/interfaces/dao"
	"github.com/D1sordxr/aviasales/src/internal/db/models"
)

type UseCase struct {
	TicketDAO ticket.TicketDAO
	OrderDAO  order.OrderDAO
}

func NewUseCase(ticketDAO ticket.TicketDAO, orderDAO order.OrderDAO) *UseCase {
	return &UseCase{
		TicketDAO: ticketDAO,
		OrderDAO:  orderDAO,
	}
}

// GetTickets TODO: model to dto
func (uc *UseCase) GetTickets() ([]models.Ticket, error) {
	return uc.TicketDAO.GetTickets()
}

func (uc *UseCase) GetTicketsDTO() (dto.Tickets, error) {
	return uc.TicketDAO.GetTicketsDTO()
}

func (uc *UseCase) GetTicketByID(id string) (models.Ticket, error) {
	return uc.TicketDAO.GetTicketByID(id)
}

func (uc *UseCase) GetTicketByIDDTO(id string) (dto.Ticket, error) {
	return uc.TicketDAO.GetTicketByIDDTO(id)
}

func (uc *UseCase) CreateTicket(t models.Ticket) error {
	return uc.TicketDAO.CreateTicket(t)
}

func (uc *UseCase) CreateTicketDTO(t dto.Ticket) (dto.Ticket, error) {
	return uc.TicketDAO.CreateTicketDTO(t)
}

func (uc *UseCase) UpdateTicket(t models.Ticket) error {
	return uc.TicketDAO.UpdateTicket(t)
}

func (uc *UseCase) UpdateTicketDTO(t dto.Ticket) (dto.Ticket, error) {
	return uc.TicketDAO.UpdateTicketDTO(t)
}

func (uc *UseCase) DeleteTicket(id string) error {
	return uc.TicketDAO.DeleteTicket(id)
}

func (uc *UseCase) DeleteTicketDTO(id string) (dto.Ticket, error) {
	return uc.TicketDAO.DeleteTicketDTO(id)
}
