package handler

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
	"ticket_wave/internal/errGroup"
	"ticket_wave/internal/models"
)

func (h Handler) CreateFeedbackHandler(c *gin.Context) {
	var req models.Feedback

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	user, _ := c.Get("user")
	req.SenderID = user.(models.Participant).ID

	_, err := h.EventService.GetEventService(models.Event{ID: req.EventID})
	if err == errGroup.EventNotFound {
		c.JSON(http.StatusNotFound, gin.H{"error": "event not found"})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	err = h.FeedbackService.CreateFeedbackService(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "feedback created"})
}

func (h Handler) GetFeedbackHandler(c *gin.Context) {
	var req models.Event

	id := c.Query("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid event id"})
		return
	}
	req.ID = idInt

	log.Println(req)

	feedbacks, err := h.FeedbackService.GetFeedbackService(req)
	log.Println(err)
	if err == errGroup.EventNotFound {
		c.JSON(http.StatusNotFound, gin.H{"error": "event not found"})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"feedback": feedbacks})
}

func (h Handler) DeleteFeedbackHandler(c *gin.Context) {
	var req models.Feedback

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	user, _ := c.Get("user")
	req.SenderID = user.(models.Participant).ID

	err := h.FeedbackService.DeleteFeedbackService(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "feedback deleted"})
}

func (h Handler) AnswerFeedbackHandler(c *gin.Context) {
	var req models.Feedback

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	log.Println(req)
	err := h.FeedbackService.AnswerFeedbackService(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "feedback answered"})
}
