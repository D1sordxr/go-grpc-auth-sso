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
