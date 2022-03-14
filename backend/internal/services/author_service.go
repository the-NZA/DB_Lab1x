package services

import (
	"github.com/the-NZA/DB_Lab1/backend/internal/models"
	"github.com/the-NZA/DB_Lab1/backend/internal/store/storer"
)

type AuthorService struct {
	repository storer.AuthorRepository
}

func (a *AuthorService) Get(ID string) (models.Author, error) {
	return a.repository.Get(ID)
}

func (a *AuthorService) Add(author models.AuthorWithBooks) (models.AuthorWithBooks, error) {
	// validate new author
	if err := author.Validate(); err != nil {
		return author, err
	}

	return a.repository.Add(author)
}

func (a *AuthorService) Update(author models.AuthorWithBooks) (models.AuthorWithBooks, error) {
	// validate updated author
	if err := author.Validate(); err != nil {
		return author, err
	}

	return a.repository.Update(author)
}

func (a *AuthorService) Delete(ID string) error {
	return a.repository.Delete(ID)
}

func (a *AuthorService) GetAll() ([]models.Author, error) {
	return a.repository.GetAll()
}
