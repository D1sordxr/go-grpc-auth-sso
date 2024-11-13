package models

import "time"

type Ticket struct {
	Base
	PassengerName string
	Payment       uint
	Destination   string
	DispatchTime  time.Time // время отправки
	ArrivalTime   time.Time // время прибытия
}
