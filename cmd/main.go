package main

import (
	"GymEventTracker/internal/database"
	"GymEventTracker/internal/features/attendance"
	"GymEventTracker/internal/features/members"
	"html/template"
	"io"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	dbPath := os.Getenv("DB_PATH")
	serverPort := os.Getenv("SERVER_PORT")

	// Initialize the SQLite database
	database.InitDB(dbPath)
	defer database.CloseDB()

	// Create a new Echo instance
	e := echo.New()
	parsedTemplates, err := template.ParseFiles("../internal/features/members/templates/index.html")

	if err != nil {
		log.Fatalf("Error parsing templates")
	}

	t := &Template{
		templates: parsedTemplates,
	}

	e.Renderer = t
	// Setup routes
	members.SetupRoutes(e)
	attendance.SetupRoutes(e)

	e.Logger.Fatal(e.Start(serverPort))
}
