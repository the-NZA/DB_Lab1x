package mock

import (
	"github.com/the-NZA/DB_Lab1/backend/internal/config"
	"github.com/the-NZA/DB_Lab1/backend/internal/store/storer"
)

type MockStore struct {
	books   *BookRepository
	genres  *GenreRepository
	authors *AuthorRepository
}

func (s *MockStore) Books() storer.BookReporsitory {
	if s.books != nil {
		return s.books
	}

	s.books = &BookRepository{}

	return s.books
}

func (s *MockStore) Authors() storer.AuthorRepository {
	if s.authors != nil {
		return s.authors
	}

	s.authors = &AuthorRepository{}

	return s.authors
}

func (s *MockStore) Genres() storer.GenreRepository {
	if s.genres != nil {
		return s.genres
	}

	s.genres = &GenreRepository{}

	return s.genres
}

func (s *MockStore) BooksAuthors() storer.BookAuthorRepository {
	return nil
}

func (s *MockStore) Close() error {
	return nil
}

func NewStore(c *config.Config) (storer.Storer, error) {
	if c == nil {
		return nil, config.ErrEmptyConfig
	}

	return &MockStore{}, nil
}
