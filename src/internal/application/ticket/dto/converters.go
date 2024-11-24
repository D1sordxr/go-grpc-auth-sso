package dto

import "github.com/D1sordxr/aviasales/src/internal/db/models"

func (t *Ticket) ToModel() models.Ticket {
	return models.Ticket{
		OrderID:       t.OrderID,
		IsAvailable:   t.IsAvailable,
		PassengerName: t.PassengerName,
		Destination:   t.Destination,
		Payment:       t.Payment,
		DispatchTime:  t.DispatchTime,
		ArrivalTime:   t.ArrivalTime,
	}
}

func (t *Ticket) ModelToDTO(ticket models.Ticket) Ticket {
	return Ticket{
		OrderID:       ticket.OrderID,
		IsAvailable:   ticket.IsAvailable,
		PassengerName: ticket.PassengerName,
		Destination:   ticket.Destination,
		Payment:       ticket.Payment,
		DispatchTime:  ticket.DispatchTime,
		ArrivalTime:   ticket.ArrivalTime,
	}
}
