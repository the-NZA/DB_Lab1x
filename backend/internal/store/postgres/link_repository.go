package postgres

import (
	"log"
	"strconv"

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

// GetByBookID find links
func (l *LinkRepository) GetByBookID(bookID string) ([]models.Link, error) {
	var links []models.Link

	// Get all links
	err := l.db.Select(&links, "SELECT * FROM links WHERE book_id = ? AND deleted != true", bookID)
	if err != nil {
		return nil, err
	}

	return links, nil
}

// Add method save new link
func (l *LinkRepository) Add(link models.Link) (models.Link, error) {
	res, err := l.db.Exec("INSERT INTO links (book_id, link) VALUES (?, ?)", link.BookID, link.Link)
	if err != nil {
		log.Println(err)
		return link, err
	}

	// Try get ID for inserted link
	id, err := res.LastInsertId()
	if err != nil {
		return link, err
	}

	// Save string representation of ID
	link.ID = strconv.FormatInt(id, 10)

	return link, nil
}

// Delete link from db
func (l *LinkRepository) Delete(linkID string) error {
	_, err := l.db.Exec("UPDATE links SET deleted = true WHERE id = ?", linkID)
	return err
}
