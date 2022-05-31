package postgres

import (
	"errors"

	"github.com/jmoiron/sqlx"
	"github.com/the-NZA/DB_Lab1x/backend/internal/models"
)

var (
	ErrEmptyIDs = errors.New("you must pass at least 1 ID")
)

type BooksAuthorsRepository struct {
	db *sqlx.DB
}

// Get all books and authors by bookID
func (b *BooksAuthorsRepository) getAll() ([]models.BookAuthor, error) {
	var booksAuthors []models.BookAuthor

	err := b.db.Select(&booksAuthors, "SELECT * FROM books_authors WHERE deleted != true")
	if err != nil {
		return nil, err
	}

	return booksAuthors, nil
}

// Get books and authors by bookID
func (b *BooksAuthorsRepository) getByBookID(bookID string) ([]models.BookAuthor, error) {
	var booksAuthors []models.BookAuthor

	err := b.db.Select(&booksAuthors, "SELECT * FROM books_authors WHERE book_id = $1 AND deleted != true", bookID)
	if err != nil {
		return nil, err
	}

	return booksAuthors, nil
}

// Get books and authors by authorID
func (b *BooksAuthorsRepository) getByAuthorID(authorID string) ([]models.BookAuthor, error) {
	var booksAuthors []models.BookAuthor

	err := b.db.Select(&booksAuthors, "SELECT * FROM books_authors WHERE author_id = $2 AND deleted != true", authorID)
	if err != nil {
		return nil, err
	}

	return booksAuthors, nil
}

// Get books and authors by authorID
func (b *BooksAuthorsRepository) getByIDs(bookID, authorID string) ([]models.BookAuthor, error) {
	var booksAuthors []models.BookAuthor

	err := b.db.Select(&booksAuthors, "SELECT * FROM books_authors WHERE book_id = $1 AND author_id = $2 AND deleted != true", bookID, authorID)
	if err != nil {
		return nil, err
	}

	return booksAuthors, nil
}

// GetByIDs get books and authors slice by passed IDs
func (b *BooksAuthorsRepository) GetByIDs(bookID, authorID string) ([]models.BookAuthor, error) {
	hasBookID := bookID != ""
	hasAuthorID := authorID != ""

	if !hasBookID && !hasAuthorID {
		return b.getAll()
	}

	if !hasAuthorID {
		return b.getByBookID(bookID)
	}

	if !hasBookID {
		return b.getByAuthorID(authorID)
	}

	return b.getByIDs(bookID, authorID)
}
