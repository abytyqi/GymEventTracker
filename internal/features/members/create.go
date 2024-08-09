package members

import (
	"GymEventTracker/internal/database"
	"GymEventTracker/internal/database/models"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

// createMember adds a new member to the database
func createMember(repo database.MemberRepo) func(c echo.Context) error {

	return func(c echo.Context) error {
		var newMember models.Member
		if err := c.Bind(&newMember); err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{
				"error": "Invalid request payload",
			})
		}

		newMember.JoinedDate = time.Now().Format("2006-01-02")
		id, err := repo.Create(newMember)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{
				"error": "Error executing statement",
			})
		}

		newMember.ID = int(id)
		return c.JSON(http.StatusCreated, newMember)
	}
}
