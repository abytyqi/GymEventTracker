package main

import (
	"GymEventTracker/internal/database"
	"GymEventTracker/internal/features/attendance"
	"GymEventTracker/internal/features/members"
	"html/template"
	"io"

	"github.com/labstack/echo/v4"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
	// Initialize the SQLite database
	database.InitDB("gym_event_tracker.db")
	defer database.CloseDB()

	// Create a new Echo instance
	e := echo.New()

	t := &Template{
		templates: template.Must(template.ParseFiles("/Users/albans.bytyci/go/src/gymeventtracker/internal/features/members/templates/index.html")),
	}

	e.Renderer = t
	// Setup routes
	members.SetupRoutes(e)
	attendance.SetupRoutes(e)

	e.Logger.Fatal(e.Start(":3000"))
}
