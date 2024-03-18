package models

import "time"

// Ticket represents a ticket purchased for an event.
type Ticket struct {
	ID           int       `json:"id"`
	EventID      int       `json:"eventId"`
	OwnerID      int       `json:"ownerId"`
	PurchaseDate time.Time `json:"purchaseDate"`
	SeatNumber   string    `json:"seatNumber"`
}
