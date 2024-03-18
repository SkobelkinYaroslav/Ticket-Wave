package repository

import (
	"database/sql"
	"ticket_wave/internal/models"
)

type EventRepo struct {
	db *sql.DB
}

func NewEventRepo(db *sql.DB) *EventRepo {
	return &EventRepo{
		db: db,
	}
}

func (e *EventRepo) CreateEvent(event models.Event) (models.Event, error) {
	query := `INSERT INTO event (organizer_id, name, description, category, date_time, address, ticket_price) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id`
	err := e.db.QueryRow(query, event.OrganizerID, event.Name, event.Description, event.Category, event.DateTime, event.Address, event.TicketPrice).Scan(&event.ID)
	if err != nil {
		return models.Event{}, err
	}
	return event, nil
}

func (e *EventRepo) GetEvent(id int) (models.Event, error) {
	var event models.Event
	query := `SELECT * FROM event WHERE id = $1`
	err := e.db.QueryRow(query, id).Scan(&event.ID, &event.OrganizerID, &event.Name, &event.Description, &event.Category, &event.DateTime, &event.Address, &event.TicketPrice)
	if err != nil {
		return models.Event{}, err
	}
	return event, nil
}

func (e *EventRepo) UpdateEvent(event models.Event) error {
	query := `UPDATE event SET organizer_id = $1, name = $2, description = $3, category = $4, date_time = $5, address = $6, ticket_price = $7 WHERE id = $8`
	_, err := e.db.Exec(query, event.OrganizerID, event.Name, event.Description, event.Category, event.DateTime, event.Address, event.TicketPrice, event.ID)
	if err != nil {
		return err
	}
	return nil
}

func (e *EventRepo) DeleteEvent(id int) error {
	query := `DELETE FROM event WHERE id = $1`
	_, err := e.db.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}
