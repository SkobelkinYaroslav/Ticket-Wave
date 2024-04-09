package models

import "database/sql"

type Feedback struct {
	ID       int            `json:"id,omitempty"`
	EventID  int            `json:"eventId"`
	SenderID int            `json:"senderId"`
	Text     string         `json:"text"`
	Reply    sql.NullString `json:"response"`
}
