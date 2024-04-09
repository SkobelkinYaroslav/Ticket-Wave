package repository

import "database/sql"

type Repository struct {
	Auth        AuthRepo
	Event       EventRepository
	Feedback    FeedbackRepo
	Participant ParticipantRepo
	Ticket      TicketRepo
}

// TODO: проверить SQL запросы
func NewRepository(db *sql.DB) Repository {
	return Repository{
		Auth:        NewAuthRepo(db),
		Event:       NewEventRepository(db),
		Feedback:    NewFeedbackRepo(db),
		Participant: NewParticipantRepo(db),
		Ticket:      NewTicketRepo(db),
	}
}
