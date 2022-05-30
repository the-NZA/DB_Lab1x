package postgres

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // Postgres driver

	"github.com/the-NZA/DB_Lab1x/backend/internal/config"
	"github.com/the-NZA/DB_Lab1x/backend/internal/store/storer"
)

type Store struct {
	db           *sqlx.DB
	books        storer.BookReporsitory
	authors      storer.AuthorRepository
	genres       storer.GenreRepository
	users        storer.UserRepository
	links        storer.LinkRepository
	booksAuthors storer.BookAuthorRepository
}

func (s *Store) Books() storer.BookReporsitory {
	if s.books == nil {
		s.books = &BookRepository{db: s.db}
	}

	return s.books
}

func (s *Store) Authors() storer.AuthorRepository {
	if s.authors == nil {
		s.authors = &AuthorRepository{db: s.db}
	}

	return s.authors
}

func (s *Store) Genres() storer.GenreRepository {
	if s.genres == nil {
		s.genres = &GenreRepository{db: s.db}
	}

	return s.genres
}

func (s *Store) Users() storer.UserRepository {
	if s.users == nil {
		s.users = &UserRepository{db: s.db}
	}

	return s.users
}

func (s *Store) Links() storer.LinkRepository {
	if s.links == nil {
		s.links = &LinkRepository{db: s.db}
	}

	return s.links
}

func (s *Store) BooksAuthors() storer.BookAuthorRepository {
	if s.booksAuthors == nil {
		s.booksAuthors = &BooksAuthorsRepository{db: s.db}
	}

	return s.booksAuthors
}

func (s *Store) Close() error {
	if s == nil {
		return storer.ErrStoreNilOrEmpty
	}

	return s.db.Close()
}

func NewStore(c *config.Config) (storer.Storer, error) {
	if c == nil {
		return nil, config.ErrEmptyConfig
	}

	dburl := fmt.Sprintf("%s", c.DBURL)

	db, err := sqlx.Connect(c.DBType, dburl)
	if err != nil {
		return nil, err
	}

	return &Store{
		db: db,
	}, nil
}
