package mock

import (
	"fmt"
	"math/rand"
	"strconv"

	"github.com/the-NZA/DB_Lab1/backend/internal/models"
)

var genres = []models.Genre{
	{
		ID:      "1",
		Title:   "First",
		Snippet: "First genre description",
		Deleted: false,
	},
	{
		ID:      "2",
		Title:   "Second",
		Snippet: "Second genre description",
		Deleted: false,
	},
	{
		ID:      "3",
		Title:   "Third",
		Snippet: "Third genre description",
		Deleted: false,
	},
}

// Get next ID
func getNextGenreID(g []models.Genre) int {
	if len(g) < 1 {
		return 0
	}

	next, err := strconv.Atoi(g[0].ID)
	if err != nil {
		return (rand.Int() % len(g)) + 1
	}

	for i := range g {
		n, err := strconv.Atoi(g[i].ID)
		if err != nil {
			continue
		}

		if n > next {
			next = n
		}
	}

	next++
	return next
}

type GenreRepository struct{}

// find one genre by passed ID
func (g GenreRepository) findByID(ID string) (int, error) {
	for i := range genres {
		if genres[i].ID == ID {
			return i, nil
		}
	}

	return -1, fmt.Errorf("Genre not found")
}

// Get one genre from genres
func (g GenreRepository) Get(ID string) (models.Genre, error) {
	idx, err := g.findByID(ID)
	if err != nil {
		return models.Genre{}, fmt.Errorf("Genre not found")
	}

	return genres[idx], nil
}

// Add one genre to genres
func (g GenreRepository) Add(genre models.Genre) (models.Genre, error) {
	newID := getNextGenreID(genres)

	genre.ID = strconv.Itoa(newID)

	genres = append(genres, genre)

	return genre, nil
}

// Update one genre in genres
func (g GenreRepository) Update(genre models.Genre) (models.Genre, error) {
	idx, err := g.findByID(genre.ID)
	if err != nil {
		return genre, fmt.Errorf("Genre not found")
	}

	genres[idx] = genre

	return genres[idx], nil
}

// Delete one genre from genres
func (g GenreRepository) Delete(ID string) error {
	idx, err := g.findByID(ID)
	if err != nil {
		return fmt.Errorf("Genre not found")
	}

	genres[idx] = genres[len(genres)-1]
	genres = genres[:len(genres)-1]

	return nil
}

// Get all genres from genres
func (g GenreRepository) GetAll() ([]models.Genre, error) {
	if len(genres) < 1 {
		return nil, fmt.Errorf("Genres not found")
	}

	return genres, nil
}
