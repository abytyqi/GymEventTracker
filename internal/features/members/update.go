package members

import (
	"GymEventTracker/internal/database"
	"GymEventTracker/internal/database/models"
	"github.com/labstack/echo/v4"
	"net/http"
)

// updateMember updates an existing member in the database
func updateMember(c echo.Context) error {
	id := c.Param("id")

	var updatedMember models.Member
	if err := c.Bind(&updatedMember); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": "Invalid request payload",
		})
	}

	stmt, err := database.DB.Prepare("UPDATE members SET name = ?, email = ?, age = ? WHERE id = ?")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": "Error preparing statement",
		})
	}
	defer stmt.Close()

	_, err = stmt.Exec(updatedMember.Name, updatedMember.Email, updatedMember.Age, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": "Error executing statement",
		})
	}

	return c.JSON(http.StatusOK, updatedMember)
}
