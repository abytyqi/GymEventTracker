package sql_lite

import (
	"GymEventTracker/internal/database"
	"GymEventTracker/internal/database/models"
)

type memberRepo struct {
}

func NewSqlLiteMemberRepo() database.MemberRepo {
	return &memberRepo{}
}

func (r memberRepo) Create(m models.Member) (int64, error) {
	stmt, err := db.Prepare("INSERT INTO members (name, email, age, joined_date) VALUES (?, ?, ?, ?)")
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

func (r memberRepo) Delete(id int64) error {
	stmt, err := db.Prepare("DELETE FROM members WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	return err
}

func (r memberRepo) Update(m models.Member) error {
	return nil
}

func (r memberRepo) Get(id int64) (models.Member, error) {
	return models.Member{}, nil
}

func (r memberRepo) List() ([]models.Member, error) {
	rows, err := db.Query("SELECT id, name, email, age, joined_date FROM members")
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
