package dto

import (
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
