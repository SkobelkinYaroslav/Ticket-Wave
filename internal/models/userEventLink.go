package models

type UserEventLink struct {
	Participant
	Event       Event  `json:"event"`
	LinkType    string `json:"linkType"`
	TicketCount int    `json:"ticketCount,omitempty"`
}
