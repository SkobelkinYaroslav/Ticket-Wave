package handler

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"ticket_wave/internal/errGroup"
	"ticket_wave/internal/models"
)

func (h Handler) CreateEventHandler(c *gin.Context) {
	var req models.Event

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	user, _ := c.Get("user")

	if user.(models.Participant).Role != "organizer" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "permission denied"})
		return
	}

	req.OrganizerID = user.(models.Participant).ID

	event, err := h.EventService.CreateEventService(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": event})
}

func (h Handler) PatchEventHandler(c *gin.Context) {
	var req models.Event

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	err := h.EventService.PatchEventService(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "event updated"})
}

func (h Handler) GetEventHandler(c *gin.Context) {
	var req models.Event

	log.Println("GetEventHandler: ", req)
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	events, err := h.EventService.GetEventService(req)
	if err == errGroup.EventNotFound {
		c.JSON(http.StatusNotFound, gin.H{"error": "event not found"})
		return
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, events)
}

func (h Handler) InteractWithEventHandler(c *gin.Context) {
	var req models.UserEventLink

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	user, ok := c.Get("user")
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user"})
		return
	}

	participant, ok := user.(models.Participant)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "User is not of type Participant"})
		return
	}

	req.Participant = participant

	_, err := h.EventService.GetEventService(req.Event)
	if err == errGroup.EventNotFound {
		c.JSON(http.StatusNotFound, gin.H{"error": "event not found"})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	err = h.EventService.InteractWithEventService(req)
	if err == errGroup.NotEnoughTickets {
		c.JSON(http.StatusConflict, gin.H{"error": "Not enough tickets"})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success"})
}
