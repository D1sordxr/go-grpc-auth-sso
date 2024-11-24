package vo

import (
	"errors"
	"github.com/D1sordxr/aviasales/src/internal/domain/ticket/exceptions"
)

type TicketPrice struct {
	Price uint64
}

// Create price is just a simple domain Value Object implementation and can be omitted here because of uint64 type
func (t *TicketPrice) Create(price uint64) (TicketPrice, error) {
	if price < 0 {
		err := exceptions.Exception{
			Error: errors.New("price can not be negative"),
		}
		return TicketPrice{}, err.Error
	}

	return TicketPrice{Price: price}, nil
}
