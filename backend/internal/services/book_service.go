package services

import (
	"github.com/the-NZA/DB_Lab1/backend/internal/models"
	"github.com/the-NZA/DB_Lab1/backend/internal/store/storer"
)

type BookService struct {
	repository storer.BookReporsitory
}

func (b *BookService) Get(ID string) (models.Book, error) {
	return b.repository.Get(ID)
}

func (b *BookService) Add(ba models.BookWithAuthors) (models.BookWithAuthors, error) {
	// validate new book with authors
	if err := ba.Validate(); err != nil {
		return ba, err
	}

	return b.repository.Add(ba)
}

func (b *BookService) Update(ba models.BookWithAuthors) (models.BookWithAuthors, error) {
	// validate updated book
	if err := ba.Validate(); err != nil {
		return ba, err
	}

	return b.repository.Update(ba)
}

func (b *BookService) Delete(ID string) error {
	return b.repository.Delete(ID)
}

func (b *BookService) GetAll() ([]models.Book, error) {
	return b.repository.GetAll()
}
