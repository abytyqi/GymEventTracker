package members

import (
	"GymEventTracker/internal/database"
	"github.com/labstack/echo/v4"
	"net/http"
)

// deleteMember deletes a member from the database
func deleteMember(c echo.Context) error {
	id := c.Param("id")

	stmt, err := database.DB.Prepare("DELETE FROM members WHERE id = ?")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": "Error preparing statement",
		})
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": "Error executing statement",
		})
	}

	return c.JSON(http.StatusNoContent, nil)
}
