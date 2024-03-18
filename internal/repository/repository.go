package repository

import (
	"database/sql"
	"ticket_wave/internal/models"
)

type Participant interface {
	CreateParticipant(participant models.Participant) (models.Participant, error)
	GetParticipant(id int) (models.Participant, error)
	UpdateParticipant(participant models.Participant) error
}

type Event interface {
	CreateEvent(event models.Event) (models.Event, error)
	GetEvent(id int) (models.Event, error)
	UpdateEvent(event models.Event) error
	DeleteEvent(id int) error
}

type EventFeedback interface {
	CreateEventFeedback(feedback models.EventFeedback) error
	GetEventFeedback(event models.Event) ([]models.EventFeedback, error)
	GetParticipantFeedback(participant models.Participant) ([]models.EventFeedback, error)
	UpdateEventFeedback(feedback models.EventFeedback) error
	DeleteEventFeedback(feedback models.EventFeedback) error
}

type Ticket interface {
	CreateTicket(ticket models.Ticket) (models.Ticket, error)
	GetTicket(participant models.Participant, event models.Event) (models.Ticket, error)
	DeleteTicket(val models.Ticket) error
}

type UserEventLink interface {
	CreateUserEventLink(link models.UserEventLink) error
	GetUserEventLink(link models.UserEventLink) (models.UserEventLink, error)
	UpdateUserEventLink(link models.UserEventLink) error
	DeleteUserEventLink(link models.UserEventLink) error
}

type Repository struct {
	Participant
	Event
	EventFeedback
	Ticket
	UserEventLink
}

func NewRepository(sql *sql.DB) *Repository {
	return &Repository{
		Participant:   nil,
		Event:         nil,
		EventFeedback: nil,
		Ticket:        nil,
		UserEventLink: nil,
	}
}
