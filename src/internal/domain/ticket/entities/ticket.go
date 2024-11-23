package entities

import "time"

type Ticket struct {
	OrderID       *int
	IsAvailable   *bool
	PassengerName *string
	Destination   *string
	Payment       *uint64
	DispatchTime  *time.Time
	ArrivalTime   *time.Time
}
