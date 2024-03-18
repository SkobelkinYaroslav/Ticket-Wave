package repository

import (
	"database/sql"
	"ticket_wave/internal/models"
)

type EventFeedbackRepo struct {
	db *sql.DB
}

func NewEventFeedbackRepo(db *sql.DB) *EventFeedbackRepo {
	return &EventFeedbackRepo{
		db: db,
	}
}

func (e *EventFeedbackRepo) CreateEventFeedback(feedback models.EventFeedback) error {
	query := `INSERT INTO event_feedback (event_id, sender_id, text) VALUES ($1, $2, $3) RETURNING id`
	err := e.db.QueryRow(query, feedback.EventID, feedback.SenderID, feedback.Text).Scan(&feedback.ID)

	if err != nil {
		return err
	}

	return nil
}

func (e *EventFeedbackRepo) GetEventFeedback(event models.Event) ([]models.EventFeedback, error) {
	var feedbacks []models.EventFeedback
	query := `SELECT * FROM event_feedback WHERE event_id = $1`
	rows, err := e.db.Query(query, event.ID)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var feedback models.EventFeedback
		err := rows.Scan(&feedback.ID, &feedback.EventID, &feedback.SenderID, &feedback.Text)
		if err != nil {
			return nil, err
		}
		feedbacks = append(feedbacks, feedback)

	}
	return feedbacks, nil
}

func (e *EventFeedbackRepo) GetParticipantFeedback(participant models.Participant) ([]models.EventFeedback, error) {
	var feedbacks []models.EventFeedback
	query := `SELECT * FROM event_feedback WHERE sender_id = $1`
	rows, err := e.db.Query(query, participant.ID)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var feedback models.EventFeedback
		err := rows.Scan(&feedback.ID, &feedback.EventID, &feedback.SenderID, &feedback.Text)
		if err != nil {
			return nil, err
		}
		feedbacks = append(feedbacks, feedback)

	}
	return feedbacks, nil
}

func (e *EventFeedbackRepo) UpdateEventFeedback(feedback models.EventFeedback) error {
	query := `UPDATE event_feedback SET text = $1 WHERE id = $2`
	_, err := e.db.Exec(query, feedback.Text, feedback.ID)
	if err != nil {
		return err
	}
	return nil
}

func (e *EventFeedbackRepo) DeleteEventFeedback(feedback models.EventFeedback) error {
	query := `DELETE FROM event_feedback WHERE id = $1`
	_, err := e.db.Exec(query, feedback.ID)
	return err
}
