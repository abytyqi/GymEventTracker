package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// Home handler function to display a welcome message
func Home(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{
		"message": "Welcome to GymEventTracker Alban!",
	})
}

// ListEvents handler function to list all events
func ListEvents(c echo.Context) error {
	events := []string{"Yoga Class", "Crossfit Session", "Zumba Dance"}
	return c.JSON(http.StatusOK, map[string][]string{
		"events": events,
	})
}

// CreateEvent handler function to create a new event
func CreateEvent(c echo.Context) error {
	// Define the Event struct to hold the event data
	type Event struct {
		Name string `json:"name" validate:"required"`
		Time string `json:"time" validate:"required"`
	}

	var newEvent Event

	// Bind and validate the incoming JSON request
	if err := c.Bind(&newEvent); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid input data",
		})
	}

	if newEvent.Name == "" || newEvent.Time == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Name and time are required",
		})
	}

	// Return a success response with the created event
	return c.JSON(http.StatusCreated, map[string]interface{}{
		"message": "Event created successfully",
		"event":   newEvent,
	})
}
