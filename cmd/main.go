package main

import (
	"GymEventTracker/internal/database"
	"GymEventTracker/internal/routes"
	"os"
	"os/signal"
	"syscall"

	"github.com/labstack/echo/v4"
)

func main() {
	// Initialize the SQLite database
	database.InitDB("gym_event_tracker.db")
	defer database.CloseDB()

	// Create a new Echo instance
	e := echo.New()

	// Setup routes
	routes.SetupRoutes(e)

	// Graceful shutdown
	go func() {
		if err := e.Start(":3000"); err != nil {
			e.Logger.Info("Shutting down the server...")
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	if err := e.Shutdown(nil); err != nil {
		e.Logger.Fatal(err)
	}
}
