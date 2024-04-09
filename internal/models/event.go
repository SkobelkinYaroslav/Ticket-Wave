package models

import "time"

type Event struct {
	ID          int       `json:"id,omitempty"`
	OrganizerID int       `json:"organizerId,omitempty"`
	Name        string    `json:"name,omitempty"`
	Description string    `json:"description,omitempty"`
	Category    string    `json:"category,omitempty"`
	DateTime    time.Time `json:"dateTime,omitempty"`
	Venue       string    `json:"venue,omitempty"`
	Address     string    `json:"address,omitempty"`
	TicketPrice float64   `json:"ticketPrice,omitempty"`
	TicketCount int       `json:"ticketCount,omitempty"`
}
