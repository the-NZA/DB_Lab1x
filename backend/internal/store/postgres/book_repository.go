package postgres

import (
	"fmt"
	"strconv"

	"github.com/jmoiron/sqlx"
	"github.com/the-NZA/DB_Lab1x/backend/internal/models"
)

var (
	insertBook = `INSERT INTO books (title, snippet, pages_cnt, pub_year, genre_id) 
			VALUES ($1, $2, $3, $4, $5) RETURNING id`
	updateBook = `UPDATE books 
			SET 	title = $1, 
				snippet = $2, 
				pages_cnt = $3, 
				pub_year = $4, 
				deleted = $5, 
				genre_id = $6  
			WHERE id = $7`
)

type BookRepository struct {
	db *sqlx.DB
}

// Get one book from books
func (b *BookRepository) Get(ID string) (models.Book, error) {
	book := models.Book{}

	err := b.db.Get(&book, "SELECT * FROM books WHERE id = $1 AND deleted != true", ID)
	if err != nil {
		return models.Book{}, err
	}

	return book, nil
}

// Add one book to books
func (b *BookRepository) Add(ba models.BookWithAuthors) (models.BookWithAuthors, error) {
	// Try save new book

	insertedID := 0

	// Start transaction
	tx, err := b.db.Beginx()
	if err != nil {
		return ba, err
	}

	// Save new book
	err = tx.QueryRowx(insertBook,
		ba.Book.Title,
		ba.Book.Snippet,
		ba.Book.PagesCnt,
		ba.Book.PublishYear,
		ba.Book.GenreID,
	).Scan(&insertedID)
	if err != nil {
		tx.Rollback()
		return ba, err
	}

	// Save string representation of ID
	ba.Book.ID = strconv.Itoa(insertedID)

	// Remove previous information about book if it exists and authors relations
	_, err = tx.Exec("DELETE FROM books_authors WHERE book_id = $1", ba.Book.ID)
	if err != nil {
		tx.Rollback()
		return ba, err
	}

	// Insert new information about book and authors relations
	for i := range ba.AuthorsIDs {
		_, err = tx.Exec("INSERT INTO books_authors (book_id, author_id) VALUES ($1, $2)", ba.Book.ID, ba.AuthorsIDs[i])
		if err != nil {
			tx.Rollback()
			return ba, err
		}
	}

	// Commit transaction
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return ba, err
	}

	return ba, nil
}

// Update one book in books
func (b *BookRepository) Update(ba models.BookWithAuthors) (models.BookWithAuthors, error) {
	// Start transaction
	tx, err := b.db.Beginx()
	if err != nil {
		return ba, err
	}

	// Remove previous information about book if it exists and authors relations
	_, err = tx.Exec("DELETE FROM books_authors WHERE book_id = $1", ba.Book.ID)
	if err != nil {
		tx.Rollback()
		return ba, err
	}

	// Insert new information about book and authors relations
	for i := range ba.AuthorsIDs {
		_, err = tx.Exec("INSERT INTO books_authors (book_id, author_id) VALUES ($1, $2)", ba.Book.ID, ba.AuthorsIDs[i])
		if err != nil {
			tx.Rollback()
			return ba, err
		}
	}

	// Try update book
	_, err = tx.Exec(updateBook,
		ba.Book.Title,
		ba.Book.Snippet,
		ba.Book.PagesCnt,
		ba.Book.PublishYear,
		ba.Book.Deleted,
		ba.Book.GenreID,
		ba.Book.ID,
	)
	if err != nil {
		tx.Rollback()
		return ba, err
	}

	// Commit transaction
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return ba, err
	}

	return ba, nil
}

// Delete one book from books
func (b *BookRepository) Delete(ID string) error {
	// Start transaction
	tx, err := b.db.Beginx()
	if err != nil {
		return err
	}

	_, err = tx.Exec("UPDATE books SET deleted = true WHERE id = $1", ID)
	if err != nil {
		tx.Rollback()
		return err
	}

	// Remove previous information about book if it exists and authors relations
	_, err = tx.Exec("DELETE FROM books_authors WHERE book_id = $1", ID)
	if err != nil {
		tx.Rollback()
		return err
	}

	// Commit transaction
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return err
	}

	return nil
}

// GetRandom3 return 3 random books from books table
func (b *BookRepository) GetRandom3() ([]models.Book, error) {
	var books []models.Book

	// Get all books from database
	err := b.db.Select(&books, "SELECT * FROM books WHERE deleted != true ORDER BY RANDOM() LIMIT 3")
	if err != nil {
		return nil, err
	}

	return books, nil
}

// GetAll return all books
func (b *BookRepository) GetAll() ([]models.Book, error) {
	var books []models.Book

	// Get all books from database
	err := b.db.Select(&books, "SELECT * FROM books WHERE deleted != true")
	if err != nil {
		return nil, err
	}

	return books, nil
}

// Search books by title, author and genre from books
func (b *BookRepository) Search(title, author, genre string) ([]models.Book, error) {
	switch {
	case len(title) > 0:
		return b.searchByTitle(title)
	case len(author) > 0:
		return b.searchByAuthor(author)
	case len(genre) > 0:
		return b.searchByGenre(genre)
	default:
		return []models.Book{}, fmt.Errorf("unsupported search param")
	}
}

func (b *BookRepository) searchByTitle(title string) ([]models.Book, error) {
	var books []models.Book

	// Get all books from database
	err := b.db.Select(&books, "SELECT * FROM books WHERE deleted != true AND title LIKE '%'||$1||'%'", title)
	if err != nil {
		return nil, err
	}

	return books, nil
}

func (b *BookRepository) searchByAuthor(author string) ([]models.Book, error) {
	var books []models.Book

	// Get all books from database
	err := b.db.Select(&books, "select b.id, b.title, b.snippet, b.pages_cnt, b.pub_year, b.deleted, b.genre_id from books b join books_authors ba on b.id = ba.book_id join authors au on ba.author_id = au.id where au.lastname like '%'||$1||'%'", author)
	if err != nil {
		return nil, err
	}

	return books, nil
}

func (b *BookRepository) searchByGenre(genre string) ([]models.Book, error) {
	var books []models.Book

	// Get all books from database
	err := b.db.Select(&books, "select b.id, b.title, b.snippet, b.pages_cnt, b.pub_year, b.deleted, b.genre_id from books b join genres g on b.genre_id = g.id where g.title like '%'||$1||'%'", genre)
	if err != nil {
		return nil, err
	}

	return books, nil
}
