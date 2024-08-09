package database

import (
	"GymEventTracker/internal/database/models"
)

type MemberRepo interface {
	Create(m models.Member) (int64, error)
	Delete(id int64) error
	Update(m models.Member) error
	Get(id int64) (models.Member, error)
	List() ([]models.Member, error)
}
