package members

import (
	"GymEventTracker/internal/database"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

// createMember adds a new member to the database
func createMember(c echo.Context) error {
	var newMember Member
	if err := c.Bind(&newMember); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": "Invalid request payload",
		})
	}

	newMember.JoinedDate = time.Now().Format("2006-01-02")

	stmt, err := database.DB.Prepare("INSERT INTO members (name, email, age, joined_date) VALUES (?, ?, ?, ?)")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": "Error preparing statement",
		})
	}
	defer stmt.Close()

	res, err := stmt.Exec(newMember.Name, newMember.Email, newMember.Age, newMember.JoinedDate)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": "Error executing statement",
		})
	}

	id, err := res.LastInsertId()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": "Error retrieving last insert ID",
		})
	}

	newMember.ID = int(id)

	return c.JSON(http.StatusCreated, newMember)
}
