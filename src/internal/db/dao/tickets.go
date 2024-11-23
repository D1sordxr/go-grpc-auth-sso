package dao

import (
	"context"
	"fmt"
	"github.com/D1sordxr/aviasales/src/internal/application/ticket/dto"
	"github.com/D1sordxr/aviasales/src/internal/db/models"
	"github.com/jackc/pgx/v5"
	"time"
)

type TicketDAO struct {
	DB *pgx.Conn
}

func NewTicketDAO(conn *pgx.Conn) *TicketDAO {
	return &TicketDAO{DB: conn}
}

func (dao *TicketDAO) GetTickets() ([]models.Ticket, error) {
	var ticket models.Ticket
	rows, err := dao.DB.Query(context.Background(), `
	SELECT id, passenger_name, destination, payment, dispatch_time, arrival_time, is_available FROM tickets
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
			&ticket.IsAvailable); err != nil {
			return nil, err
		}
		data = append(data, ticket)
	}

	return data, nil
}

func (dao *TicketDAO) GetTicketsDTO() (dto.Tickets, error) {
	var ticket dto.Ticket
	tickets := make([]dto.Ticket, 0, 20)

	rows, err := dao.DB.Query(context.Background(), `
	SELECT passenger_name, destination, payment, dispatch_time, arrival_time, is_available, order_id FROM tickets
	 `)
	if err != nil {
		return dto.Tickets{}, err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(
			&ticket.PassengerName,
			&ticket.Destination,
			&ticket.Payment,
			&ticket.DispatchTime,
			&ticket.ArrivalTime,
			&ticket.IsAvailable,
			&ticket.OrderID,
		)
		if err != nil {
			return dto.Tickets{}, err
		}

		tickets = append(tickets, ticket)
	}

	return dto.Tickets{Tickets: tickets}, nil
}

func (dao *TicketDAO) GetTicketByID(id string) (models.Ticket, error) {
	var ticket models.Ticket

	err := dao.DB.QueryRow(context.Background(), `
		SELECT passenger_name, destination, payment, dispatch_time, arrival_time, is_available 
		FROM tickets WHERE ID = $1
	`, id).Scan(
		&ticket.PassengerName,
		&ticket.Destination,
		&ticket.Payment,
		&ticket.DispatchTime,
		&ticket.ArrivalTime,
		&ticket.IsAvailable,
	)
	if err != nil {
		return models.Ticket{}, err
	}

	return ticket, nil
}

func (dao *TicketDAO) GetTicketByIDDTO(id string) (dto.Ticket, error) {
	var ticket dto.Ticket

	err := dao.DB.QueryRow(context.Background(), `
		SELECT passenger_name, destination, payment, dispatch_time, arrival_time, is_available, order_id 
		FROM tickets WHERE ID = $1
	`, id).Scan(
		&ticket.PassengerName,
		&ticket.Destination,
		&ticket.Payment,
		&ticket.DispatchTime,
		&ticket.ArrivalTime,
		&ticket.IsAvailable,
		&ticket.OrderID,
	)
	if err != nil {
		return dto.Ticket{}, err
	}

	return ticket, nil
}

func (dao *TicketDAO) CreateTicket(t models.Ticket) error {
	_, err := dao.DB.Exec(context.Background(), `
		INSERT INTO tickets(passenger_name, destination, payment, dispatch_time, arrival_time, created_at) 
		VALUES ($1, $2, $3, $4, $5, NOW())`,
		t.PassengerName, t.Destination, t.Payment, t.DispatchTime, t.ArrivalTime,
	)
	if err != nil {
		return err
	}

	return nil
}

func (dao *TicketDAO) CreateTicketDTO(t dto.Ticket) (dto.Ticket, error) {
	var ticket dto.Ticket
	err := dao.DB.QueryRow(context.Background(), `
		INSERT INTO tickets(passenger_name, destination, payment, dispatch_time, arrival_time, created_at, order_id) 
		VALUES ($1, $2, $3, $4, $5, NOW(), $6)
		RETURNING passenger_name, destination, payment, dispatch_time, arrival_time, is_available, order_id`,
		t.PassengerName, t.Destination, t.Payment, t.DispatchTime, t.ArrivalTime, nil,
	).Scan(
		&ticket.PassengerName,
		&ticket.Destination,
		&ticket.Payment,
		&ticket.DispatchTime,
		&ticket.ArrivalTime,
		&ticket.IsAvailable,
		&ticket.OrderID,
	)
	if err != nil {
		return dto.Ticket{}, err
	}

	return ticket, nil
}

func (dao *TicketDAO) UpdateTicket(t models.Ticket) error {
	query :=
		`UPDATE tickets
	SET passenger_name = COALESCE(NULLIF($1::TEXT, ''), passenger_name),
	destination = COALESCE(NULLIF($2::TEXT, ''), destination),
	payment = COALESCE($3::INTEGER, payment),
	dispatch_time = COALESCE($4::TIMESTAMP, dispatch_time),
	arrival_time = COALESCE($5::TIMESTAMP, arrival_time),
	updated_at = NOW()
	WHERE id = $6::INTEGER`

	passengerName := ""
	if t.PassengerName != nil {
		passengerName = *t.PassengerName
	}

	destination := ""
	if t.Destination != nil {
		destination = *t.Destination
	}

	var payment *uint64
	if t.Payment != nil {
		payment = t.Payment
	}

	var dispatchTime *time.Time
	if t.DispatchTime != nil {
		dispatchTime = t.DispatchTime
	}

	var arrivalTime *time.Time
	if t.ArrivalTime != nil {
		arrivalTime = t.ArrivalTime
	}

	_, err := dao.DB.Exec(context.Background(), query,
		passengerName, destination, payment,
		dispatchTime, arrivalTime, t.ID,
	)

	if err != nil {
		return fmt.Errorf("failed to update ticket: %w", err)
	}

	return nil
}

func (dao *TicketDAO) DeleteTicket(id string) error {
	_, err := dao.DB.Exec(context.Background(), `
		DELETE FROM tickets WHERE ID = $1
	`, id)
	if err != nil {
		return err
	}

	return nil
}

func (dao *TicketDAO) DeleteTicketDTO(id string) (dto.Ticket, error) {
	ticket, err := dao.GetTicketByIDDTO(id)
	if err != nil {
		return dto.Ticket{}, err
	}

	_, err = dao.DB.Exec(context.Background(), `
		DELETE FROM tickets WHERE ID = $1
	`, id)
	if err != nil {
		return dto.Ticket{}, err
	}

	return ticket, err
}
