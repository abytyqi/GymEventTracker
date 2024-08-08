package routes

import (
	"GymEventTracker/internal/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	r.GET("/", handlers.Home)
	r.GET("/events", handlers.ListEvents)
	r.POST("/events", handlers.CreateEvent)
}
