package members

import (
	"GymEventTracker/internal/database"
	"database/sql"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

// showMember retrieves a single member by ID
func showMember(repo database.MemberRepo) func(c echo.Context) error {
	return func(c echo.Context) error {
		idParam := c.Param("id")
		id, err := strconv.ParseInt(idParam, 10, 64)
		if err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{
				"error": "Invalid member id",
			})
		}

		member, err := repo.Get(id)
		if err != nil {
			if err == sql.ErrNoRows {
				return c.JSON(http.StatusNotFound, echo.Map{
					"error": "Member not found",
				})
			}
			return c.JSON(http.StatusInternalServerError, echo.Map{
				"error": "Error retrieving member",
			})
		}

		return c.JSON(http.StatusOK, member)
	}
}
