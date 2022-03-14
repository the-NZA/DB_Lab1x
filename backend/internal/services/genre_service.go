package services

import (
	"github.com/go-pkgz/lgr"
	"github.com/the-NZA/DB_Lab1/backend/internal/models"
	"github.com/the-NZA/DB_Lab1/backend/internal/store/storer"
)

type GenreService struct {
	repository storer.GenreRepository
}

func (g *GenreService) Get(ID string) (models.Genre, error) {
	return g.repository.Get(ID)
}

func (g *GenreService) Add(genre models.Genre) (models.Genre, error) {
	//validate new genre
	if err := genre.Validate(); err != nil {
		return genre, err
	}

	return g.repository.Add(genre)
}

func (g *GenreService) Update(genre models.Genre) (models.Genre, error) {
	lgr.Printf("DEBUG %v\n", genre)
	//validate updated genre
	if err := genre.Validate(); err != nil {
		return genre, err
	}
	return g.repository.Update(genre)
}

func (g *GenreService) Delete(ID string) error {
	return g.repository.Delete(ID)
}

func (g *GenreService) GetAll() ([]models.Genre, error) {
	return g.repository.GetAll()
}
