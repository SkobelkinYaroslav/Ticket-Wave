package service

import "ticket_wave/internal/models"

func (s Service) CreateTicketService(req models.Ticket) (models.Ticket, error) {
	return s.TicketRepository.CreateTicketRepo(req)
}

func (s Service) GetTicketService(participant models.Participant) ([]models.Ticket, error) {
	return s.TicketRepository.GetTicketRepo(participant)
}
