package main

import (
	"GymEventTracker/internal/database/sql_lite"
	"GymEventTracker/internal/features/members"
	"GymEventTracker/internal/features/users"
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
	sql_lite.InitDB(dbPath)
	defer sql_lite.CloseDB()

	// Create a new Echo instance
	e := echo.New()
	// Add Logger middleware
	e.Static("/plugins", "../plugins")
	e.Static("/dist", "../dist")
	parsedTemplates, err := template.ParseGlob("../internal/templates/*")

	if err != nil {
		log.Printf("%v", err)
		log.Fatalf("Error parsing templates")
	}

	t := &Template{
		templates: parsedTemplates,
	}

	e.Renderer = t
	// Setup routes

	//http.Handle("/plugins/", http.StripPrefix("/plugins/", http.FileServer(http.Dir("/GymEventTracker/plugins"))))

	mr := sql_lite.NewSqlLiteMemberRepo()
	ur := sql_lite.NewSqlLiteUserRepo()

	members.SetupRoutes(e, mr)
	users.SetupRoutes(e, ur)

	e.Logger.Fatal(e.Start(serverPort))
}
