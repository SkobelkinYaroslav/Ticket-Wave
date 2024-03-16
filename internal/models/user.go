package models

import "time"

type Participant struct {
	ID             int       `json:"id"`
	FirstName      string    `json:"firstName"`
	LastName       string    `json:"lastName"`
	Email          string    `json:"email"`
	Phone          string    `json:"phone,omitempty"` // omitempty to allow empty values
	BirthDate      time.Time `json:"birthDate,omitempty"`
	Username       string    `json:"username"`
	Password       string    `json:"password"`
	Role           string    `json:"role"` // Could be "user" or "organizer"
	Preferences    string    `json:"preferences,omitempty"`
	FavoriteEvents string    `json:"favoriteEvents,omitempty"`
}
