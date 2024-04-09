package service

import "ticket_wave/internal/models"

func (s Service) CreateFeedbackService(feedback models.Feedback) error {
	return s.FeedbackRepository.CreateFeedbackRepo(feedback)
}

func (s Service) GetFeedbackService(event models.Event) ([]models.Feedback, error) {
	return s.FeedbackRepository.GetFeedbackRepo(event)
}

func (s Service) DeleteFeedbackService(feedback models.Feedback) error {
	return s.FeedbackRepository.DeleteFeedbackRepo(feedback)
}

func (s Service) AnswerFeedbackService(feedback models.Feedback) error {
	return s.FeedbackRepository.AnswerFeedbackRepo(feedback)
}
