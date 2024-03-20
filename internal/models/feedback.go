package models

// EventFeedback represents feedback left by users for an event.
type EventFeedback struct {
	ID       int    `json:"id"`
	EventID  int    `json:"eventId"`
	SenderID int    `json:"senderId"`
	Text     string `json:"text"`
}
