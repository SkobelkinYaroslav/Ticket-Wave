package models

import "time"

type Ticket struct {
	ID           int       `json:"id,omitempty"`
	EventID      int       `json:"eventId"`
	OwnerID      int       `json:"ownerId"`
	PurchaseDate time.Time `json:"purchaseDate"`
	SeatNumber   string    `json:"seatNumber"`
}
