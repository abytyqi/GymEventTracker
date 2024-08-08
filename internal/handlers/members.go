package handlers

import (
	"GymEventTracker/internal/database"
	"database/sql"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

// Member represents a gym member
type Member struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Email      string `json:"email"`
	Age        int    `json:"age"`
	JoinedDate string `json:"joined_date"`
}

// GetMembers retrieves all members from the database
func GetMembers(c echo.Context) error {
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

// CreateMember adds a new member to the database
func CreateMember(c echo.Context) error {
	var newMember Member
	if err := c.Bind(&newMember); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": "Invalid request payload",
		})
	}

	newMember.JoinedDate = time.Now().Format("2006-01-02")

	stmt, err := database.DB.Prepare("INSERT INTO members (name, email, age, joined_date) VALUES (?, ?, ?, ?)")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": "Error preparing statement",
		})
	}
	defer stmt.Close()

	res, err := stmt.Exec(newMember.Name, newMember.Email, newMember.Age, newMember.JoinedDate)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": "Error executing statement",
		})
	}

	id, err := res.LastInsertId()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": "Error retrieving last insert ID",
		})
	}

	newMember.ID = int(id)

	return c.JSON(http.StatusCreated, newMember)
}

// GetMember retrieves a single member by ID
func GetMember(c echo.Context) error {
	id := c.Param("id")

	var member Member
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

// UpdateMember updates an existing member in the database
func UpdateMember(c echo.Context) error {
	id := c.Param("id")

	var updatedMember Member
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

// DeleteMember deletes a member from the database
func DeleteMember(c echo.Context) error {
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
