package repository

import (
	"database/sql"
	"ticket_wave/internal/errGroup"
	"ticket_wave/internal/models"
)

type EventRepository struct {
	db *sql.DB
}

func NewEventRepository(db *sql.DB) EventRepository {
	return EventRepository{
		db: db,
	}
}

func (e EventRepository) CreateEventRepo(event models.Event) (models.Event, error) {
	query := `INSERT INTO event (organizer_id, name, description, category, date_time, address, ticket_price, ticket_count) VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id`
	err := e.db.QueryRow(query, event.OrganizerID, event.Name, event.Description, event.Category, event.DateTime, event.Address, event.TicketPrice, event.TicketCount).Scan(&event.ID)
	if err != nil {
		return models.Event{}, err
	}
	return event, nil
}
func (e EventRepository) GetEventRepo(req models.Event) (models.Event, error) {
	var event models.Event
	query := `SELECT * FROM event WHERE id = $1`
	err := e.db.QueryRow(query, req.ID).Scan(&event.ID, &event.OrganizerID, &event.Name, &event.Description, &event.Category, &event.DateTime, &event.Address, &event.TicketPrice, &event.TicketCount)
	if err == sql.ErrNoRows {
		return models.Event{}, errGroup.EventNotFound
	}

	if err != nil {
		return models.Event{}, err
	}
	return event, nil
}

func (e EventRepository) GetAllEventsRepo() ([]models.Event, error) {
	var events []models.Event
	query := `SELECT * FROM event`
	rows, err := e.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var event models.Event
		err := rows.Scan(&event.ID, &event.OrganizerID, &event.Name, &event.Description, &event.Category, &event.DateTime, &event.Address, &event.TicketPrice, &event.TicketCount)
		if err != nil {
			return nil, err
		}
		events = append(events, event)
	}
	return events, nil

}

func (e EventRepository) PatchEventRepo(event models.Event) error {
	query := `UPDATE event SET name = $1, description = $2, category = $3, date_time = $4, address = $5, ticket_price = $6, ticket_count = $7 WHERE id = $8`
	_, err := e.db.Exec(query, event.Name, event.Description, event.Category, event.DateTime, event.Address, event.TicketPrice, event.TicketCount, event.ID)
	if err != nil {
		return err
	}
	return nil
}

func (e EventRepository) InteractWithEvent(userEventLink models.UserEventLink) error {
	query := `
		INSERT INTO user_event_link (user_id, event_id, link_type, ticket_count) 
		VALUES ($1, $2, $3, $4)
		ON CONFLICT (user_id, event_id) 
		DO UPDATE SET link_type = $3, ticket_count = $4`
	_, err := e.db.Exec(query, userEventLink.Participant.ID, userEventLink.Event.ID, userEventLink.LinkType, userEventLink.TicketCount)
	if err != nil {
		return err
	}

	query = `UPDATE event SET ticket_count = ticket_count - $1 WHERE id = $2`
	_, err = e.db.Exec(query, userEventLink.TicketCount, userEventLink.Event.ID)
	if err != nil {
		return err
	}

	return nil
}
