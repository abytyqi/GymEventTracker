package routes

import (
	"GymEventTracker/internal/handlers"

	"github.com/labstack/echo/v4"
)

func SetupRoutes(e *echo.Echo) {
	e.GET("/", handlers.Home)
	e.GET("/events", handlers.ListEvents)
	e.POST("/events", handlers.CreateEvent)
}
