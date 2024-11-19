package dto

import "time"

type Ticket struct {
	OrderID       *int       `json:"order_id"`
	IsAvailable   *bool      `json:"is_available"`
	PassengerName *string    `json:"passenger_name"`
	Destination   *string    `json:"destination"`
	Payment       *uint64    `json:"payment"`       // price
	DispatchTime  *time.Time `json:"dispatch_time"` // время отправки
	ArrivalTime   *time.Time `json:"arrival_time"`  // время прибытия
}

type Tickets struct {
	Tickets []Ticket `json:"tickets"`
}
