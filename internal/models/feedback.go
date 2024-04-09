package models

type Feedback struct {
	ID       int    `json:"id,omitempty"`
	EventID  int    `json:"eventId"`
	SenderID int    `json:"senderId"`
	Text     string `json:"text"`
	Reply    string `json:"response"`
}
