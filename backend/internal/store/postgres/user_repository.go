package postgres

import (
	"strconv"

	"github.com/jmoiron/sqlx"
	"github.com/the-NZA/DB_Lab1x/backend/internal/models"
)

var (
	insertUser        = `INSERT INTO users (username, passwd, email) VALUES (?, ?, ?)`
	getUserByUsername = `SELECT * FROM users WHERE username = ? AND deleted != true`
)

type UserRepository struct {
	db *sqlx.DB
}

// Find user
func (u *UserRepository) GetByUsername(username string) (models.User, error) {
	user := models.User{}

	err := u.db.Get(&user, getUserByUsername, username)
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

// Save new user
func (u *UserRepository) Add(user models.User) (models.User, error) {
	res, err := u.db.Exec(insertUser, user.Username, user.HashedPassword, user.Email)
	if err != nil {
		return user, err
	}

	// Try get inserted ID
	id, err := res.LastInsertId()
	if err != nil {
		return user, err
	}

	// Save string representation of ID
	user.ID = strconv.FormatInt(id, 10)

	return user, nil
}
