package members

import (
	"GymEventTracker/internal/database"
	"github.com/labstack/echo/v4"
	"net/http"
)

// getMembers retrieves all members from the database
func getMembers(c echo.Context) error {
	rows, err := database.DB.Query("SELECT id, name, email, age, joined_date FROM members")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": "Unable to retrieve members",
		})
	}
	defer rows.Close()

	var members []Member
	for rows.Next() {
		var member Member
		err := rows.Scan(&member.ID, &member.Name, &member.Email, &member.Age, &member.JoinedDate)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{
				"error": "Error scanning member data",
			})
		}
		members = append(members, member)
	}

	return c.JSON(http.StatusOK, members)
}
