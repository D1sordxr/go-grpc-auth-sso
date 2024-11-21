package dto

import (
	"github.com/D1sordxr/aviasales/src/internal/db/models"
	"time"
)

type Ticket struct {
	OrderID       *int       `json:"order_id,omitempty"`
	IsAvailable   *bool      `json:"is_available,omitempty"`
	PassengerName *string    `json:"passenger_name,omitempty"`
	Destination   *string    `json:"destination,omitempty"`
	Payment       *uint64    `json:"payment,omitempty"`       // price
	DispatchTime  *time.Time `json:"dispatch_time,omitempty"` // время отправки
	ArrivalTime   *time.Time `json:"arrival_time,omitempty"`  // время прибытия
}

type Tickets struct {
	Tickets []Ticket `json:"tickets"`
}

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
