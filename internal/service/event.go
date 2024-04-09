package service

import (
	"ticket_wave/internal/errGroup"
	"ticket_wave/internal/models"
)

func (s Service) CreateEventService(event models.Event) (models.Event, error) {
	return s.EventRepository.CreateEventRepo(event)
}

func (s Service) PatchEventService(event models.Event) error {
	return s.EventRepository.PatchEventRepo(event)
}

func (s Service) GetEventService(event models.Event) (models.Event, error) {
	return s.EventRepository.GetEventRepo(event)
}

func (s Service) GetAllEventsService() ([]models.Event, error) {
	return s.EventRepository.GetAllEventsRepo()
}

func (s Service) InteractWithEventService(link models.UserEventLink) error {
	event, err := s.GetEventService(link.Event)

	if err == errGroup.EventNotFound {
		return errGroup.EventNotFound
	}

	if err != nil {
		return err
	}

	if link.LinkType == "buy" && event.TicketCount < link.TicketCount {
		return errGroup.NotEnoughTickets
	}

	return s.EventRepository.InteractWithEvent(link)
}
