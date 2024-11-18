package models

import "github.com/google/uuid"

type Order struct {
	Base
	ClientID      *uuid.UUID `json:"client_id"`
	AddressID     *uuid.UUID `json:"address_id"`
	OrderStatus   *string    `json:"order_status"`
	PaymentMethod *string    `json:"payment_method"`
	SerialNumber  *int       `json:"serial_number"`
	Closed        *bool      `json:"closed"`
	Deleted       *bool      `json:"deleted"`
	Tickets       []Ticket   `json:"tickets"`
}
