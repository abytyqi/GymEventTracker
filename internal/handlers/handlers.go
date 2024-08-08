package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Home(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Welcome to GymEventTracker!",
	})
}

func ListEvents(c *gin.Context) {
	// This will list all events
	events := []string{"Yoga Class", "Crossfit Session", "Zumba Dance"}
	c.JSON(http.StatusOK, gin.H{
		"events": events,
	})
}

func CreateEvent(c *gin.Context) {
	// This will create a new event
	type Event struct {
		Name string `json:"name"`
		Time string `json:"time"`
	}

	var newEvent Event
	if err := c.ShouldBindJSON(&newEvent); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Event created successfully",
		"event":   newEvent,
	})
}
