package repository

import (
	"database/sql"
	"ticket_wave/internal/models"
)

type TicketRepo struct {
	db *sql.DB
}

func NewTicketRepo(db *sql.DB) TicketRepo {
	return TicketRepo{
		db: db,
	}
}

func (t TicketRepo) CreateTicketRepo(ticket models.Ticket) (models.Ticket, error) {
	query := `INSERT INTO ticket (event_id, owner_id, purchase_date, seat_number) VALUES ($1, $2, $3, $4) RETURNING id`
	err := t.db.QueryRow(query, ticket.EventID, ticket.OwnerID, ticket.PurchaseDate, ticket.SeatNumber).Scan(&ticket.ID)
	if err != nil {
		return models.Ticket{}, err
	}
	return ticket, nil
}

func (t TicketRepo) GetTicketRepo(participant models.Participant) ([]models.Ticket, error) {
	query := `SELECT * FROM ticket WHERE owner_id = $1`
	rows, err := t.db.Query(query, participant.ID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tickets []models.Ticket
	for rows.Next() {
		var ticket models.Ticket
		err := rows.Scan(&ticket.ID, &ticket.EventID, &ticket.OwnerID, &ticket.PurchaseDate, &ticket.SeatNumber)
		if err != nil {
			return nil, err
		}
		tickets = append(tickets, ticket)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return tickets, nil
}
