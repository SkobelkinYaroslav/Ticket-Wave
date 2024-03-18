package models

type UserEventLink struct {
	UserID   int    `json:"userId"`
	EventID  int    `json:"eventId"`
	LinkType string `json:"linkType"`
}
