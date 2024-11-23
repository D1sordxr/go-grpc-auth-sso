package entities

import (
	"github.com/D1sordxr/aviasales/src/internal/domain/ticket/vo"
	"time"
)

type Ticket struct {
	OrderID       *int
	IsAvailable   *bool
	PassengerName *string
	Destination   *string
	Payment       *vo.TicketPrice
	DispatchTime  *time.Time
	ArrivalTime   *time.Time
}
