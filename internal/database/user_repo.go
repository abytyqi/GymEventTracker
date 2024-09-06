package database

import "GymEventTracker/internal/database/models"

type UserRepo interface {
	Create(m models.User) (int64, error)
	List() ([]models.User, error)
}
