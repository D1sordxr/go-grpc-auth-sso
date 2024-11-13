package db

import (
	"aviasales/src/internal/db/models"
	"context"
)

func (s *Storage) GetTicket() ([]models.Ticket, error) {
	var ticket models.Ticket
	rows, err := s.DB.Query(context.Background(), `
	SELECT id, passenger_name, destination, payment, dispatch_time, arrival_time FROM ticket
	 `)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	data := make([]models.Ticket, 0, 10)
	for rows.Next() {
		if err = rows.Scan(
			&ticket.ID, &ticket.PassengerName, &ticket.Destination,
			&ticket.Payment, &ticket.DispatchTime, &ticket.ArrivalTime,
		); err != nil {
			return nil, err
		}
		data = append(data, ticket)
	}

	return data, nil
}

func (s *Storage) CreateTicket(t models.Ticket) error {
	_, err := s.DB.Exec(context.Background(), `
		INSERT INTO ticket(passenger_name, destination, payment, dispatch_time, arrival_time) 
		VALUES ($1, $2, $3, $4, $5)`,
		t.PassengerName, t.Destination, t.Payment, t.DispatchTime, t.ArrivalTime,
	)
	if err != nil {
		return err
	}

	return nil
}
