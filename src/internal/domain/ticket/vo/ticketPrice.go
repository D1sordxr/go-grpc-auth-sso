package vo

import (
	"errors"
	"github.com/D1sordxr/aviasales/src/internal/domain/ticket/exceptions"
)

type TicketPrice struct {
	Price int
}

func (t *TicketPrice) Create(price int) (TicketPrice, error) {
	if price < 0 {
		err := exceptions.Exception{
			Error: errors.New("price can not be negative"),
		}
		return TicketPrice{}, err.Error
	}

	return TicketPrice{Price: price}, nil
}
