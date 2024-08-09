package members

import (
	"GymEventTracker/internal/database"
	"github.com/labstack/echo/v4"
	"net/http"
)

// getMembers retrieves all members from the database
func getMembers(repo database.MemberRepo) func(c echo.Context) error {
	return func(c echo.Context) error {

		members, err := repo.List()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{
				"error": "Unable to retrieve members",
			})
		}

		return c.Render(http.StatusOK, "index", members)
	}
}
