package models

import "time"

// Ticket represents a ticket purchased for an event.
type Ticket struct {
	ID           int       `json:"id"`
	EventID      int       `json:"eventId"` // Foreign key to the Event table
	OwnerID      int       `json:"ownerId"` // Foreign key to the Participant table
	PurchaseDate time.Time `json:"purchaseDate"`
	SeatNumber   string    `json:"seatNumber"`
}
