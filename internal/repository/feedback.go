package repository

import (
	"database/sql"
	"ticket_wave/internal/models"
)

type FeedbackRepo struct {
	db *sql.DB
}

func NewFeedbackRepo(db *sql.DB) FeedbackRepo {
	return FeedbackRepo{
		db: db,
	}
}

func (e FeedbackRepo) CreateFeedbackRepo(req models.Feedback) error {
	query := `INSERT INTO event_feedback (event_id, sender_id, text) VALUES ($1, $2, $3) RETURNING id`
	err := e.db.QueryRow(query, req.EventID, req.SenderID, req.Text).Scan(&req.ID)

	if err != nil {
		return err
	}

	return nil
}

func (e FeedbackRepo) GetFeedbackRepo(req models.Event) ([]models.Feedback, error) {
	var feedbacks []models.Feedback
	query := `SELECT * FROM event_feedback WHERE event_id = $1`
	rows, err := e.db.Query(query, req.ID)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var feedback models.Feedback
		err := rows.Scan(&feedback.ID, &feedback.EventID, &feedback.SenderID, &feedback.Text, &feedback.Reply)
		if err != nil {
			return nil, err
		}
		feedbacks = append(feedbacks, feedback)

	}
	return feedbacks, nil
}

func (e FeedbackRepo) DeleteFeedbackRepo(req models.Feedback) error {
	query := `DELETE FROM event_feedback WHERE event_id = $1 AND sender_id = $2`
	_, err := e.db.Exec(query, req.EventID, req.SenderID)
	return err
}

func (e FeedbackRepo) AnswerFeedbackRepo(req models.Feedback) error {
	query := `UPDATE event_feedback SET reply = $1 WHERE id = $2 AND sender_id = $3`
	_, err := e.db.Exec(query, req.Reply.String, req.ID, req.SenderID)
	return err
}
