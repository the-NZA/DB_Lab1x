package services

import (
	"github.com/the-NZA/DB_Lab1x/backend/internal/models"
	"github.com/the-NZA/DB_Lab1x/backend/internal/store/storer"
)

type BookAuthorService struct {
	repository storer.BookAuthorRepository
}

func (ba *BookAuthorService) GetByIDs(bookID, authorID string) ([]models.BookAuthor, error) {
	return ba.repository.GetByIDs(bookID, authorID)
}
