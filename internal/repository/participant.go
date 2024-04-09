package repository

import (
	"database/sql"
	"ticket_wave/internal/models"
)

type ParticipantRepo struct {
	db *sql.DB
}

func NewParticipantRepo(db *sql.DB) ParticipantRepo {
	return ParticipantRepo{
		db: db,
	}
}

func (p *ParticipantRepo) CreateParticipant(participant models.Participant) (models.Participant, error) {
	query := `INSERT INTO participant (first_name, last_name, email, birth_date, username, password, role) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id`
	err := p.db.QueryRow(query, participant.FirstName, participant.LastName, participant.Email, participant.Password, participant.Role).Scan(&participant.ID)
	if err != nil {
		return models.Participant{}, err
	}
	return participant, nil
}
func (p *ParticipantRepo) GetParticipant(id int) (models.Participant, error) {
	var participant models.Participant

	query := `SELECT * FROM participant WHERE id = $1`

	err := p.db.QueryRow(query, id).Scan(&participant.ID, &participant.FirstName, &participant.LastName, &participant.Email, &participant.Password, &participant.Role)

	if err != nil {
		return models.Participant{}, err
	}

	return participant, nil

}

func (p *ParticipantRepo) UpdateParticipant(participant models.Participant) (models.Participant, error) {
	query := `UPDATE participant SET first_name = $1, last_name = $2, email = $3, birth_date = $4, username = $5, password = $6, role = $7 WHERE id = $8 RETURNING id`
	err := p.db.QueryRow(query, participant.FirstName, participant.LastName, participant.Email, participant.Password, participant.Role, participant.ID).Scan(&participant.ID)
	if err != nil {
		return models.Participant{}, err
	}
	return participant, nil
}
