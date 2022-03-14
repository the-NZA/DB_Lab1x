package mysql

import (
	"strconv"

	"github.com/jmoiron/sqlx"
	"github.com/the-NZA/DB_Lab1/backend/internal/models"
)

var (
	insertGenre = `INSERT INTO genres (title, snippet) VALUES (?, ?)`
	updateGenre = `UPDATE genres 
			SET 	title = ?, 
				snippet = ?, 
				deleted = ? 
			WHERE id = ?`
)

type GenreRepository struct {
	db *sqlx.DB
}

// Get one genre
func (g *GenreRepository) Get(ID string) (models.Genre, error) {
	genre := models.Genre{}

	// Try get one genre
	err := g.db.Get(&genre, "SELECT * FROM genres WHERE id = ? AND deleted != true", ID)
	if err != nil {
		return models.Genre{}, err
	}

	return genre, nil
}

// Add new genre
func (g *GenreRepository) Add(genre models.Genre) (models.Genre, error) {
	// Try save new genre
	res, err := g.db.Exec(insertGenre, genre.Title, genre.Snippet)
	if err != nil {
		return genre, err
	}

	// Try get ID for inserted genre
	id, err := res.LastInsertId()
	if err != nil {
		return genre, err
	}

	// Save string representation of ID
	genre.ID = strconv.FormatInt(id, 10)

	return genre, nil
}

// Update genre
func (g *GenreRepository) Update(genre models.Genre) (models.Genre, error) {
	// Try update genre
	_, err := g.db.Exec(updateGenre,
		genre.Title,
		genre.Snippet,
		genre.Deleted,
		genre.ID,
	)
	if err != nil {
		return genre, err
	}

	return genre, nil
}

// Delete genre
func (g *GenreRepository) Delete(ID string) error {
	_, err := g.db.Exec("UPDATE genres SET deleted = true WHERE id = ?", ID)

	return err
}

// Get all genres
func (g *GenreRepository) GetAll() ([]models.Genre, error) {
	var genres []models.Genre

	// Get all genres from database
	err := g.db.Select(&genres, "SELECT * FROM genres WHERE deleted != true")
	if err != nil {
		return nil, err
	}

	return genres, nil
}
