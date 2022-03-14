package services

import (
	"github.com/the-NZA/DB_Lab1/backend/internal/config"
	"github.com/the-NZA/DB_Lab1/backend/internal/store/storer"
)

type Services struct {
	store       storer.Storer
	config      *config.Config
	books       BookServicer
	genres      GenreServicer
	authors     AuthorServicer
	bookAuthors BooksAuthorsServicer
}

func (s *Services) BookService() BookServicer {
	if s.books != nil {
		return s.books
	}

	s.books = &BookService{
		repository: s.store.Books(),
	}

	return s.books
}

func (s *Services) AuthorService() AuthorServicer {
	if s.authors != nil {
		return s.authors
	}

	s.authors = &AuthorService{
		repository: s.store.Authors(),
	}

	return s.authors
}

func (s *Services) GenreService() GenreServicer {
	if s.genres != nil {
		return s.genres
	}

	s.genres = &GenreService{
		repository: s.store.Genres(),
	}

	return s.genres
}

func (s *Services) BooksAuthors() BooksAuthorsServicer {
	if s.bookAuthors == nil {
		s.bookAuthors = &BookAuthorService{
			repository: s.store.BooksAuthors(),
		}
	}

	return s.bookAuthors
}

func (s *Services) Close() error {
	return s.store.Close()
}

// NewServices create initiale new services structure
func NewServices(c *config.Config, s storer.Storer) (Servicer, error) {
	if c == nil {
		return nil, config.ErrEmptyConfig
	}

	if s == nil {
		return nil, storer.ErrStoreNilOrEmpty
	}

	return &Services{
		store:  s,
		config: c,
	}, nil
}
