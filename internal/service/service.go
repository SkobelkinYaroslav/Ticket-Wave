package service

import (
	"ticket_wave/internal/models"
	"ticket_wave/internal/repository"
)

type AuthRepository interface {
	CreateUserRepo(req models.Participant) error
	GetUserRepo(req models.Participant) (models.Participant, error)
}

type EventRepository interface {
	GetEventRepo(req models.Event) (models.Event, error)
	GetAllEventsRepo() ([]models.Event, error)
	CreateEventRepo(event models.Event) (models.Event, error)
	PatchEventRepo(event models.Event) error
	InteractWithEvent(userEventLink models.UserEventLink) error
}

type TicketRepository interface {
	CreateTicketRepo(ticket models.Ticket) (models.Ticket, error)
	GetTicketRepo(participant models.Participant) ([]models.Ticket, error)
}

type FeedbackRepository interface {
	CreateFeedbackRepo(req models.Feedback) error
	GetFeedbackRepo(req models.Event) ([]models.Feedback, error)
	DeleteFeedbackRepo(req models.Feedback) error
	AnswerFeedbackRepo(req models.Feedback) error
}

type Service struct {
	AuthRepository     AuthRepository
	EventRepository    EventRepository
	TicketRepository   TicketRepository
	FeedbackRepository FeedbackRepository
}

func NewService(repo repository.Repository) Service {
	return Service{
		AuthRepository:     repo.Auth,
		EventRepository:    repo.Event,
		TicketRepository:   repo.Ticket,
		FeedbackRepository: repo.Feedback,
	}
}
