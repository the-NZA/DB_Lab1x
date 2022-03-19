package mysql

import (
	"github.com/jmoiron/sqlx"
	"github.com/the-NZA/DB_Lab1x/backend/internal/models"
)

var (
	insertLink      = `INSERT INTO users (username, passwd, email) VALUES (?, ?, ?)`
	getLinkByBookID = `SELECT * FROM users WHERE username = ? AND deleted != true`
)

type LinkRepository struct {
	db *sqlx.DB
}

// Find links
func (l *LinkRepository) GetByBookID(bookID string) ([]models.Link, error) {
	// user := models.User{}

	// err := l.db.Get(&user, getUserByUsername, username)
	// if err != nil {
	// 	return models.User{}, err
	// }

	// return user, nil

	return []models.Link{}, nil
}

// Save new link
func (l *LinkRepository) Add(link models.Link) (models.Link, error) {
	// res, err := u.db.Exec(insertUser, user.Username, user.HashedPassword, user.Email)
	// if err != nil {
	// 	return user, err
	// }

	// // Try get inserted ID
	// id, err := res.LastInsertId()
	// if err != nil {
	// 	return user, err
	// }

	// // Save string representation of ID
	// user.ID = strconv.FormatInt(id, 10)

	// return user, nil
	return link, nil
}

func (l *LinkRepository) Delete(linkID string) error {
	return nil
}
