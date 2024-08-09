package members

import (
	"GymEventTracker/internal/database"
	"GymEventTracker/internal/database/models"
	"database/sql"
	"github.com/labstack/echo/v4"
	"net/http"
)

// showMember retrieves a single member by ID
func showMember(c echo.Context) error {
	id := c.Param("id")

	var member models.Member
	err := database.DB.QueryRow("SELECT id, name, email, age, joined_date FROM members WHERE id = ?", id).
		Scan(&member.ID, &member.Name, &member.Email, &member.Age, &member.JoinedDate)

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
