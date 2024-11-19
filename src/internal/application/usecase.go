package application

import (
	order "github.com/D1sordxr/aviasales/src/internal/application/order/interfaces/dao"
	ticket "github.com/D1sordxr/aviasales/src/internal/application/ticket/interfaces/dao"
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
