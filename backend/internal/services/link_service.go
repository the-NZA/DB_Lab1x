package services

import (
	"github.com/the-NZA/DB_Lab1x/backend/internal/models"
	"github.com/the-NZA/DB_Lab1x/backend/internal/store/storer"
)

type LinkService struct {
	repository storer.LinkRepository
}

func (l *LinkService) Get(bookID string) ([]models.Link, error) {
	return l.repository.GetByBookID(bookID)
}

func (l *LinkService) Add(link models.Link) (models.Link, error) {
	if err := link.Validate(); err != nil {
		return link, err
	}

	return l.repository.Add(link)
}

func (l *LinkService) Delete(linkID string) error {
	return l.repository.Delete(linkID)
}
