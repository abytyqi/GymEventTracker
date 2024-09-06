package members

import (
	"GymEventTracker/internal/database"
	"GymEventTracker/internal/database/models"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

// updateMember updates an existing member in the database
func updateMember(repo database.MemberRepo) func(c echo.Context) error {
	return func(c echo.Context) error {
		idParam := c.Param("id")
		id, err := strconv.ParseInt(idParam, 10, 64)
		if err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{
				"error": "Invalid member id",
			})
		}

		var updatedMember models.Member
		if err := c.Bind(&updatedMember); err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{
				"error": "Invalid request payload",
			})
		}
		updatedMember.ID = id

		err = repo.Update(updatedMember)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{
				"error": "Error executing statement",
			})
		}

		return c.JSON(http.StatusOK, updatedMember)
	}
}
