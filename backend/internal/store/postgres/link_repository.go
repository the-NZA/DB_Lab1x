package postgres

import (
	"log"
	"strconv"

	"github.com/jmoiron/sqlx"
	"github.com/the-NZA/DB_Lab1x/backend/internal/models"
)

type LinkRepository struct {
	db *sqlx.DB
}

// GetByBookID find links
func (l *LinkRepository) GetByBookID(bookID string) ([]models.Link, error) {
	var links []models.Link

	// Get all links
	err := l.db.Select(&links, "SELECT * FROM links WHERE book_id = $1 AND deleted != true", bookID)
	if err != nil {
		return nil, err
	}

	return links, nil
}

// Add method save new link
func (l *LinkRepository) Add(link models.Link) (models.Link, error) {
	insertedID := 0
	err := l.db.QueryRowx("INSERT INTO links (book_id, link) VALUES ($1, $2) RETURNING id", link.BookID, link.Link).Scan(&insertedID)
	if err != nil {
		log.Println(err)
		return link, err
	}

	// Save string representation of ID
	link.ID = strconv.Itoa(insertedID)

	return link, nil
}

// Delete link from db
func (l *LinkRepository) Delete(linkID string) error {
	_, err := l.db.Exec("UPDATE links SET deleted = true WHERE id = $1", linkID)
	return err
}
