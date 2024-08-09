package attendance

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

// create handler function to create a new event
func create(c echo.Context) error {
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
