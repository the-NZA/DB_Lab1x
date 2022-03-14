package mysql

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"

	"github.com/the-NZA/DB_Lab1/backend/internal/config"
	"github.com/the-NZA/DB_Lab1/backend/internal/store/storer"
)

type MySQLStore struct {
	db           *sqlx.DB
	books        storer.BookReporsitory
	authors      storer.AuthorRepository
	genres       storer.GenreRepository
	booksAuthors storer.BookAuthorRepository
}

func (s *MySQLStore) Books() storer.BookReporsitory {
	if s.books == nil {
		s.books = &BookRepository{db: s.db}
	}

	return s.books
}

func (s *MySQLStore) Authors() storer.AuthorRepository {
	if s.authors == nil {
		s.authors = &AuthorRepository{db: s.db}
	}

	return s.authors
}

func (s *MySQLStore) Genres() storer.GenreRepository {
	if s.genres == nil {
		s.genres = &GenreRepository{db: s.db}
	}

	return s.genres
}

func (s *MySQLStore) BooksAuthors() storer.BookAuthorRepository {
	if s.booksAuthors == nil {
		s.booksAuthors = &BooksAuthorsRepository{db: s.db}
	}

	return s.booksAuthors
}

func (s *MySQLStore) Close() error {
	if s == nil {
		return storer.ErrStoreNilOrEmpty
	}

	return s.db.Close()
}

func NewStore(c *config.Config) (storer.Storer, error) {
	if c == nil {
		return nil, config.ErrEmptyConfig
	}

	db, err := sqlx.Connect(c.DBType, c.DBURL)
	if err != nil {
		return nil, err
	}

	return &MySQLStore{
		db: db,
	}, nil
}
