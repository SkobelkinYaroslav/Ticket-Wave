package repository

import (
	"database/sql"
	"ticket_wave/internal/models"
)

type Participant interface {
	CreateParticipant(Participant models.Participant) (models.Participant, error)
	GetParticipant(val interface{}) (models.Participant, error)
	UpdateParticipant(Participant models.Participant) (models.Participant, error)
}

type Event interface {
	CreateEvent(event models.Event) (models.Event, error)
	GetEvent(val interface{}) (models.Event, error)
	UpdateEvent(event models.Event) (models.Event, error)
	DeleteEvent(val interface{}) error
}

type EventFeedback interface {
	CreateEventFeedback(comment string, participant models.Participant, event models.Event) error
	GetEventFeedback(val interface{}) (models.EventFeedback, error)
	UpdateEventFeedback(eventFeedback models.EventFeedback) (models.EventFeedback, error)
	DeleteEventFeedback(feedback models.EventFeedback) error
}

type Ticket interface {
	CreateTicket(ticket models.Ticket) (models.Ticket, error)
	GetTicket(val interface{}) (models.Ticket, error)
	DeleteTicket(val models.Ticket) error
}

type Repository struct {
	Participant
	Event
	EventFeedback
	Ticket
}

func NewRepository(sql *sql.DB) *Repository {
	return &Repository{
		Participant:   nil,
		Event:         nil,
		EventFeedback: nil,
		Ticket:        nil,
	}
}
