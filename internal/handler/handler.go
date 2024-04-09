package handler

import (
	"github.com/gin-gonic/gin"
	"ticket_wave/internal/models"
	"ticket_wave/internal/service"
)

type AuthService interface {
	RegisterService(req models.Participant) error
	LoginService(req models.Participant) (string, error)
	CheckTokenService(token string) (models.Participant, error)
}

type EventService interface {
	GetEventService(req models.Event) (models.Event, error)
	GetAllEventsService() ([]models.Event, error)
	CreateEventService(req models.Event) (models.Event, error)
	PatchEventService(req models.Event) error
	InteractWithEventService(link models.UserEventLink) error
}

type TicketService interface {
	CreateTicketService(req models.Ticket) (models.Ticket, error)
	GetTicketService(req models.Participant) ([]models.Ticket, error)
}

type FeedbackService interface {
	CreateFeedbackService(req models.Feedback) error
	GetFeedbackService(req models.Event) ([]models.Feedback, error)
	DeleteFeedbackService(req models.Feedback) error
	AnswerFeedbackService(req models.Feedback) error
}

type Handler struct {
	AuthService     AuthService
	EventService    EventService
	TicketService   TicketService
	FeedbackService FeedbackService
}

func NewHandler(service service.Service) *gin.Engine {
	g := gin.Default()

	handler := &Handler{
		AuthService:     service,
		EventService:    service,
		TicketService:   service,
		FeedbackService: service,
	}

	// AuthService endpoints
	g.POST("/register", handler.RegisterHandler)
	g.POST("/login", handler.LoginHandler)

	authApiGroup := g.Group("/").Use(handler.RequireAuth)
	{
		authApiGroup.POST("/event", handler.CreateEventHandler)
		authApiGroup.PATCH("/event", handler.PatchEventHandler)
		authApiGroup.GET("/event", handler.GetEventHandler)
		authApiGroup.POST("/interact", handler.InteractWithEventHandler)

		authApiGroup.POST("/ticket", handler.CreateTicketHandler)
		authApiGroup.GET("/ticket", handler.GetTicketHandler)

		authApiGroup.POST("/feedback", handler.CreateFeedbackHandler)
		authApiGroup.GET("/feedback", handler.GetFeedbackHandler)
		authApiGroup.DELETE("/feedback", handler.DeleteFeedbackHandler)
		authApiGroup.POST("/feedback/answer", handler.AnswerFeedbackHandler)
	}

	return g

}
