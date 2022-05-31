package postgres

import (
	"strconv"

	"github.com/jmoiron/sqlx"
	"github.com/the-NZA/DB_Lab1x/backend/internal/models"
)

var (
	insertGenre = `INSERT INTO genres (title, snippet) VALUES ($1, $2) RETURNING id`
	updateGenre = `UPDATE genres 
			SET 	title = $1, 
				snippet = $2, 
				deleted = $3 
			WHERE id = $4`
)

type GenreRepository struct {
	db *sqlx.DB
}

// Get one genre
func (g *GenreRepository) Get(ID string) (models.Genre, error) {
	genre := models.Genre{}

	// Try get one genre
	err := g.db.Get(&genre, "SELECT * FROM genres WHERE id = $1 AND deleted != true", ID)
	if err != nil {
		return models.Genre{}, err
	}

	return genre, nil
}

// Add new genre
func (g *GenreRepository) Add(genre models.Genre) (models.Genre, error) {
	// Try save new genre
	insertedID := 0

	err := g.db.QueryRowx(insertGenre, genre.Title, genre.Snippet).Scan(&insertedID)
	if err != nil {
		return genre, err
	}

	// Save string representation of ID
	genre.ID = strconv.Itoa(insertedID)

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
	_, err := g.db.Exec("UPDATE genres SET deleted = true WHERE id = $1", ID)

	return err
}

// GetAll returns all genres
func (g *GenreRepository) GetAll() ([]models.Genre, error) {
	var genres []models.Genre

	// Get all genres from database
	err := g.db.Select(&genres, "SELECT * FROM genres WHERE deleted != true")
	if err != nil {
		return nil, err
	}

	return genres, nil
}
