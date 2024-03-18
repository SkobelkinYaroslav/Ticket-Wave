package repository

import (
	"database/sql"
	"ticket_wave/internal/models"
)

type UserEventLinkRepo struct {
	db *sql.DB
}

func NewUserEventLinkRepo(db *sql.DB) *UserEventLinkRepo {
	return &UserEventLinkRepo{
		db: db,
	}
}

func (u *UserEventLinkRepo) CreateUserEventLink(link models.UserEventLink) error {
	query := `INSERT INTO user_event_link (user_id, event_id, link_type) VALUES ($1, $2, $3)`
	_, err := u.db.Exec(query, link.UserID, link.EventID, link.LinkType)
	return err
}

func (u *UserEventLinkRepo) GetUserEventLink(link models.UserEventLink) (models.UserEventLink, error) {
	query := `SELECT * FROM user_event_link WHERE user_id = $1 AND event_id = $2`
	err := u.db.QueryRow(query, link.UserID, link.EventID).Scan(&link.UserID, &link.EventID, &link.LinkType)
	if err != nil {
		return models.UserEventLink{}, err
	}
	return link, nil
}

func (u *UserEventLinkRepo) UpdateUserEventLink(link models.UserEventLink) error {
	query := `UPDATE user_event_link SET link_type = $1 WHERE user_id = $2 AND event_id = $3`
	_, err := u.db.Exec(query, link.LinkType, link.UserID, link.EventID)
	return err
}

func (u *UserEventLinkRepo) DeleteUserEventLink(link models.UserEventLink) error {
	query := `DELETE FROM user_event_link WHERE user_id = $1 AND event_id = $2`
	_, err := u.db.Exec(query, link.UserID, link.EventID)
	return err
}
