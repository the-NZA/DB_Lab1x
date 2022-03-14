package mysql

import (
	"strconv"

	"github.com/jmoiron/sqlx"
	"github.com/the-NZA/DB_Lab1/backend/internal/models"
)

var (
	insertAuthor = `INSERT INTO authors (firstname, lastname, surname,  snippet) 
				VALUES (?, ?, ?, ?)`
	updateAuthor = `UPDATE authors
			SET 	firstname = ?, 
				lastname = ?,
				surname = ?,
				snippet = ?, 
				deleted = ? 
			WHERE id = ?`
)

type AuthorRepository struct {
	db *sqlx.DB
}

// Get one author from db
func (a *AuthorRepository) Get(ID string) (models.Author, error) {
	author := models.Author{}

	err := a.db.Get(&author, "SELECT * FROM authors WHERE id = ? AND deleted != true", ID)
	if err != nil {
		return models.Author{}, err
	}

	return author, nil
}

// Add one author
func (a *AuthorRepository) Add(awb models.AuthorWithBooks) (models.AuthorWithBooks, error) {
	// Try save new author

	// Start transaction
	tx, err := a.db.Beginx()
	if err != nil {
		return awb, err
	}

	// Save new author
	res, err := tx.Exec(insertAuthor,
		awb.Author.Firstname,
		awb.Author.Lastname,
		awb.Author.Surname,
		awb.Author.Snippet,
	)
	if err != nil {
		tx.Rollback()
		return awb, err
	}

	// Try get ID for inserted author
	id, err := res.LastInsertId()
	if err != nil {
		tx.Rollback()
		return awb, err
	}

	// Save string representation of ID
	awb.Author.ID = strconv.FormatInt(id, 10)

	// If passed any book id for new author
	if len(awb.BooksIDs) > 0 {
		// Remove previous information about author if it exists and books relations
		_, err = tx.Exec("DELETE FROM books_authors WHERE author_id = ?", awb.Author.ID)
		if err != nil {
			tx.Rollback()
			return awb, err
		}

		// Insert information about author and books relations
		for i := range awb.BooksIDs {
			_, err = tx.Exec("INSERT INTO books_authors (book_id, author_id) VALUES (?, ?)", awb.BooksIDs[i], awb.Author.ID)
			if err != nil {
				tx.Rollback()
				return awb, err
			}
		}
	}

	// Commit transaction
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return awb, err
	}

	return awb, nil
}

// Update one author
func (a *AuthorRepository) Update(awb models.AuthorWithBooks) (models.AuthorWithBooks, error) {
	// Try update author

	// Start transaction
	tx, err := a.db.Beginx()
	if err != nil {
		return awb, err
	}

	// Remove previous information about author if it exists and books relations
	_, err = tx.Exec("DELETE FROM books_authors WHERE author_id = ?", awb.Author.ID)
	if err != nil {
		tx.Rollback()
		return awb, err
	}

	// If passed any book id for new author
	if len(awb.BooksIDs) > 0 {
		// Remove previous information about author if it exists and books relations
		_, err = tx.Exec("DELETE FROM books_authors WHERE author_id = ?", awb.Author.ID)
		if err != nil {
			tx.Rollback()
			return awb, err
		}

		// Insert information about author and books relations
		for i := range awb.BooksIDs {
			_, err = tx.Exec("INSERT INTO books_authors (book_id, author_id) VALUES (?, ?)", awb.BooksIDs[i], awb.Author.ID)
			if err != nil {
				tx.Rollback()
				return awb, err
			}
		}
	}

	// Try update author
	_, err = tx.Exec(updateAuthor,
		awb.Author.Firstname,
		awb.Author.Lastname,
		awb.Author.Surname,
		awb.Author.Snippet,
		awb.Author.Deleted,
		awb.Author.ID,
	)
	if err != nil {
		tx.Rollback()
		return awb, err
	}

	// Commit transaction
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return awb, err
	}

	return awb, nil
}

// Delete one author
func (a *AuthorRepository) Delete(ID string) error {
	// Start transaction
	tx, err := a.db.Beginx()
	if err != nil {
		return err
	}

	_, err = tx.Exec("UPDATE authors SET deleted = true WHERE id = ?", ID)
	if err != nil {
		tx.Rollback()
		return err
	}

	// Remove previous information about author if it exists and books relations
	_, err = tx.Exec("DELETE FROM books_authors WHERE author_id = ?", ID)
	if err != nil {
		tx.Rollback()
		return err
	}

	// Commit transaction
	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

// Gell all authors
func (a *AuthorRepository) GetAll() ([]models.Author, error) {
	var authors []models.Author

	// Get all authors from database
	err := a.db.Select(&authors, "SELECT * FROM authors WHERE deleted != true")
	if err != nil {
		return nil, err
	}

	return authors, nil
}
