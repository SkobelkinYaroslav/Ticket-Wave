package models

import "time"

type Event struct {
	ID          int       `json:"id"`
	EventCode   string    `json:"eventCode"`
	OrganizerID int       `json:"organizerId"`
	Name        string    `json:"name"`
	Description string    `json:"description,omitempty"`
	Category    string    `json:"category,omitempty"`
	DateTime    time.Time `json:"dateTime"`
	Venue       string    `json:"venue,omitempty"`
	Address     string    `json:"address,omitempty"`
	TicketPrice float64   `json:"ticketPrice"`
}
