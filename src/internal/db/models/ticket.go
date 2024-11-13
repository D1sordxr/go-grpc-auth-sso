package models

import "time"

type Ticket struct {
	Base
	PassengerName string
	Destination   string
	Payment       uint
	DispatchTime  time.Time // время отправки
	ArrivalTime   time.Time // время прибытия
}
