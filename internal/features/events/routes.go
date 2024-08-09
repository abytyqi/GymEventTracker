package events

import (
	"github.com/labstack/echo/v4"
)

func SetupRoutes(e *echo.Echo) {
	e.GET("/", home)
	e.GET("/events", listEvents)
	e.POST("/events", createEvent)

}
