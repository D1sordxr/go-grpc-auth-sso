package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetTicket(c *gin.Context) {
	c.JSON(http.StatusTeapot, gin.H{"message": "No ticket for you today."})
}

func CreateTicket(c *gin.Context) {
	c.JSON(http.StatusTeapot, gin.H{"message": "No NEW ticket for you today."})
}
