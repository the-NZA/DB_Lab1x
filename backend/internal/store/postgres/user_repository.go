package postgres

import (
	"strconv"

	"github.com/jmoiron/sqlx"
	"github.com/the-NZA/DB_Lab1x/backend/internal/models"
)

var (
	insertUser        = `INSERT INTO users (username, passwd, email) VALUES ($1, $2, $3) RETURNING id`
	getUserByUsername = `SELECT * FROM users WHERE username = $1 AND deleted = false`
)

type UserRepository struct {
	db *sqlx.DB
}

// GetByUsername find user
func (u *UserRepository) GetByUsername(username string) (models.User, error) {
	user := models.User{}

	err := u.db.Get(&user, getUserByUsername, username)
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

// Add handler save new user
func (u *UserRepository) Add(user models.User) (models.User, error) {
	insertedID := 0

	err := u.db.QueryRowx(insertUser, user.Username, user.HashedPassword, user.Email).Scan(&insertedID)
	if err != nil {
		return user, err
	}

	user.ID = strconv.Itoa(insertedID)

	return user, nil
}
