package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"ticket_wave/internal/errGroup"
	"ticket_wave/internal/models"
)

func (h Handler) CreateTicketHandler(c *gin.Context) {
	var req models.Ticket

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	_, err := h.EventService.GetEventService(models.Event{ID: req.EventID})
	if err == errGroup.EventNotFound {
		c.JSON(http.StatusNotFound, gin.H{"error": "event not found"})
		return
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	user, _ := c.Get("user")

	req.OwnerID = user.(models.Participant).ID

	ticket, err := h.TicketService.CreateTicketService(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": ticket})
}

func (h Handler) GetTicketHandler(c *gin.Context) {
	var req models.Participant

	user, _ := c.Get("user")

	req.ID = user.(models.Participant).ID

	tickets, err := h.TicketService.GetTicketService(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"tickets": tickets})
}
