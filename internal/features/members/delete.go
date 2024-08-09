package members

import (
	"GymEventTracker/internal/database"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

// deleteMember deletes a member from the database
func deleteMember(repo database.MemberRepo) func(c echo.Context) error {
	return func(c echo.Context) error {
		idParam := c.Param("id")
		id, err := strconv.ParseInt(idParam, 10, 64)
		if err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{
				"error": "id must be a valid number",
			})
		}

		err = repo.Delete(id)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{
				"error": "Error executing statement",
			})
		}

		return c.JSON(http.StatusNoContent, nil)
	}
}
