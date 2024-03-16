package models

// EventFeedback represents feedback left by users for an event.
type EventFeedback struct {
	ID       int    `json:"id"`
	EventID  int    `json:"eventId"`  // Foreign key to the Event table
	SenderID int    `json:"senderId"` // Foreign key to the Participant table
	Text     string `json:"text"`
}
