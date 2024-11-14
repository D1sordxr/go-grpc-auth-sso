package models

import "time"

type Ticket struct {
	Base
	OrderID       *int       `json:"order_id"`
	PassengerName *string    `json:"passenger_name"`
	Destination   *string    `json:"destination"`
	Payment       *uint64    `json:"payment"`       // price
	DispatchTime  *time.Time `json:"dispatch_time"` // время отправки
	ArrivalTime   *time.Time `json:"arrival_time"`  // время прибытия
}
