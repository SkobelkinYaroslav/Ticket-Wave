package repository

import (
	"database/sql"
	"ticket_wave/internal/models"
)

type TicketRepo struct {
	db *sql.DB
}

func NewTicketRepo(db *sql.DB) *TicketRepo {
	return &TicketRepo{
		db: db,
	}
}

func (t *TicketRepo) CreateTicket(ticket models.Ticket) (models.Ticket, error) {
	query := `INSERT INTO ticket (event_id, owner_id, purchase_date, seat_number) VALUES ($1, $2, $3, $4) RETURNING id`
	err := t.db.QueryRow(query, ticket.EventID, ticket.OwnerID, ticket.PurchaseDate, ticket.SeatNumber).Scan(&ticket.ID)
	if err != nil {
		return models.Ticket{}, err
	}
	return ticket, nil
}

func (t *TicketRepo) GetTicket(participant models.Participant, event models.Event) (models.Ticket, error) {
	var ticket models.Ticket
	query := `SELECT * FROM ticket WHERE owner_id = $1 AND event_id = $2`
	err := t.db.QueryRow(query, participant.ID, event.ID).Scan(&ticket.ID, &ticket.EventID, &ticket.OwnerID, &ticket.PurchaseDate, &ticket.SeatNumber)
	if err != nil {
		return models.Ticket{}, err
	}
	return ticket, nil
}

func (t *TicketRepo) DeleteTicket(ticket models.Ticket) error {
	query := `DELETE FROM ticket WHERE id = $1`
	_, err := t.db.Exec(query, ticket.ID)

	if err != nil {
		return err
	}

	return nil
}
