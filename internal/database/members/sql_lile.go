package members

import (
	"GymEventTracker/internal/database"
	"GymEventTracker/internal/database/models"
)

type memberSqlLiteRepo struct {
}

func NewSqlLiteMemberRepo() database.MemberRepo {
	return &memberSqlLiteRepo{}
}

func (r memberSqlLiteRepo) Create(m models.Member) (int64, error) {
	stmt, err := database.DB.Prepare("INSERT INTO members (name, email, age, joined_date) VALUES (?, ?, ?, ?)")
	if err != nil {
		return -1, err
	}
	defer stmt.Close()

	res, err := stmt.Exec(m.Name, m.Email, m.Age, m.JoinedDate)
	if err != nil {
		return -1, err
	}

	return res.LastInsertId()
}

func (r memberSqlLiteRepo) Delete(id int64) error {
	stmt, err := database.DB.Prepare("DELETE FROM members WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	return err
}

func (r memberSqlLiteRepo) Update(m models.Member) error {
	return nil
}

func (r memberSqlLiteRepo) Get(id int64) (models.Member, error) {
	return models.Member{}, nil
}

func (r memberSqlLiteRepo) List() ([]models.Member, error) {
	rows, err := database.DB.Query("SELECT id, name, email, age, joined_date FROM members")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var members []models.Member
	for rows.Next() {
		var member models.Member
		err := rows.Scan(&member.ID, &member.Name, &member.Email, &member.Age, &member.JoinedDate)
		if err != nil {
			return nil, err
		}
		members = append(members, member)
	}

	return members, nil
}
