package attendance

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

// home handler function to display a welcome message
func home(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{
		"message": "Welcome to GymEventTracker Alban!",
	})
}
