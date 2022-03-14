package mysql

import (
	"errors"

	"github.com/jmoiron/sqlx"
	"github.com/the-NZA/DB_Lab1/backend/internal/models"
)

var (
	ErrEmptyIDs = errors.New("You must pass at least 1 ID")
)

var (
	insertBooksAuthorssAuthors = `INSERT INTO books (title, snippet, pages_cnt, pub_year, genre_id) 
			VALUES (?, ?, ?, ?, ?)`
	updateBooksAuthorssAuthors = `UPDATE books 
			SET 	title = ?, 
				snippet = ?, 
				pages_cnt = ?, 
				pub_year = ?, 
				deleted = ?, 
				genre_id = ? 
			WHERE id = ?`
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

	err := b.db.Select(&booksAuthors, "SELECT * FROM books_authors WHERE book_id = ? AND deleted != true", bookID)
	if err != nil {
		return nil, err
	}

	return booksAuthors, nil
}

// Get books and authors by authorID
func (b *BooksAuthorsRepository) getByAuthorID(authorID string) ([]models.BookAuthor, error) {
	var booksAuthors []models.BookAuthor

	err := b.db.Select(&booksAuthors, "SELECT * FROM books_authors WHERE author_id = ? AND deleted != true", authorID)
	if err != nil {
		return nil, err
	}

	return booksAuthors, nil
}

// Get books and authors by authorID
func (b *BooksAuthorsRepository) getByIDs(bookID, authorID string) ([]models.BookAuthor, error) {
	var booksAuthors []models.BookAuthor

	err := b.db.Select(&booksAuthors, "SELECT * FROM books_authors WHERE book_id = ? AND author_id = ? AND deleted != true", bookID, authorID)
	if err != nil {
		return nil, err
	}

	return booksAuthors, nil
}

// Get books and authors slice by passed IDs
func (b *BooksAuthorsRepository) GetByIDs(bookID, authorID string) ([]models.BookAuthor, error) {
	hasBookID := len(bookID) != 0
	hasAuthorID := len(authorID) != 0

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
