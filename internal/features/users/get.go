package users

import (
	"GymEventTracker/internal/database"
	"github.com/labstack/echo/v4"
	"net/http"
)

func getUsers(repo database.UserRepo) func(c echo.Context) error {
	return func(c echo.Context) error {

		users, err := repo.List()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{
				"error": "Unable to retrieve users",
			})
		}

		//return c.Render(http.StatusOK, "index", users)
		return c.JSON(http.StatusOK, users)
	}
}
