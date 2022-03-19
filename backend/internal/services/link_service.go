package services

import (
	"github.com/the-NZA/DB_Lab1x/backend/internal/models"
	"github.com/the-NZA/DB_Lab1x/backend/internal/store/storer"
)

type LinkService struct {
	repository storer.LinkRepository
}

func (l *LinkService) Get(bookID string) ([]models.Link, error) {
	links, err := l.repository.GetByBookID(bookID)
	if err != nil {
		return links, err
	}

	return links, nil
}

func (l *LinkService) Add(link models.Link) (models.Link, error) {
	return l.repository.Add(link)
}

func (l *LinkService) Delete(linkID string) error {
	return l.repository.Delete(linkID)
}
