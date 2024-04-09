package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"ticket_wave/internal/models"
)

func (h Handler) RegisterHandler(c *gin.Context) {
	var req models.Participant

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	fmt.Println(req)

	if err := h.AuthService.RegisterService(req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	fmt.Println("user registered")

	c.JSON(http.StatusOK, gin.H{"message": "user registered"})
}

func (h Handler) LoginHandler(c *gin.Context) {
	var req models.Participant

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	log.Println(req)

	token, err := h.AuthService.LoginService(req)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	log.Println(token)

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", token, 3600*24*30, "", "", false, true)

	c.JSON(http.StatusOK, gin.H{"message": "user logged in"})

}

func (h Handler) RequireAuth(c *gin.Context) {
	tokenString, err := c.Cookie("Authorization")

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	user, err := h.AuthService.CheckTokenService(tokenString)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	c.Set("user", user)

	c.Next()
}
