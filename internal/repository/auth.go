package repository

import (
	"database/sql"
	"ticket_wave/internal/models"
)

type AuthRepo struct {
	db *sql.DB
}

func NewAuthRepo(db *sql.DB) AuthRepo {
	return AuthRepo{db: db}
}

func (r AuthRepo) CreateUserRepo(req models.Participant) error {
	query := `INSERT INTO participant (first_name, last_name, email, password, role) VALUES ($1, $2, $3, $4, $5) RETURNING id`
	_, err := r.db.Exec(query, req.FirstName, req.LastName, req.Email, req.Password, req.Role)
	if err != nil {
		return err
	}
	return nil
}
func (r AuthRepo) GetUserRepo(req models.Participant) (models.Participant, error) {
	var participant models.Participant
	query := `SELECT * FROM participant WHERE email = $1`
	err := r.db.QueryRow(query, req.Email).Scan(&participant.ID, &participant.FirstName, &participant.LastName, &participant.Email, &participant.Password, &participant.Role)

	if err != nil {
		return models.Participant{}, err
	}
	return participant, nil
}
