package mysql

import (
	"github.com/jmoiron/sqlx"
	"github.com/the-NZA/DB_Lab1x/backend/internal/models"
)

type UserRepository struct {
	db *sqlx.DB
}

func (u *UserRepository) GetByUsername(username string) (models.User, error) {
	return models.User{}, nil
}

func (u *UserRepository) Add(models.User) (models.User, error) {
	return models.User{}, nil
}
