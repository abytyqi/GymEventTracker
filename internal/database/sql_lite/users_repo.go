package sql_lite

import (
	"GymEventTracker/internal/database"
	"GymEventTracker/internal/database/models"
)

type userRepo struct {
}

func NewSqlLiteUserRepo() database.UserRepo {
	return &userRepo{}
}

func (r userRepo) Create(u models.User) (int64, error) {
	stmt, err := db.Prepare("INSERT INTO users (email, password) VALUES (?, ?)")
	if err != nil {
		return -1, err
	}
	defer stmt.Close()

	res, err := stmt.Exec(u.Email, u.Password)
	if err != nil {
		return -1, err
	}

	return res.LastInsertId()
}

func (r userRepo) List() ([]models.User, error) {
	rows, err := db.Query("SELECT id, email FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := make([]models.User, 0)
	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.ID, &user.Email)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}
