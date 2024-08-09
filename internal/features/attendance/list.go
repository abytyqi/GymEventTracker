package attendance

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

// list handler function to list all attendance
func list(c echo.Context) error {
	events := []string{"Yoga Class", "Crossfit Session", "Zumba Dance"}
	return c.JSON(http.StatusOK, map[string][]string{
		"attendance": events,
	})
}
