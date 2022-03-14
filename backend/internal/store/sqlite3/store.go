package sqlite3

// "db_url": "file:sqlite_data/lab.db?foreign_keys=on",

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"

	"github.com/the-NZA/DB_Lab1/backend/internal/config"
	"github.com/the-NZA/DB_Lab1/backend/internal/store/storer"
)

type SQLiteStore struct {
	db           *sqlx.DB
	books        storer.BookReporsitory
	authors      storer.AuthorRepository
	genres       storer.GenreRepository
	booksAuthors storer.BookAuthorRepository
}

func (s *SQLiteStore) Books() storer.BookReporsitory {
	if s.books == nil {
		s.books = &BookRepository{db: s.db}
	}

	return s.books
}

func (s *SQLiteStore) Authors() storer.AuthorRepository {
	if s.authors == nil {
		s.authors = &AuthorRepository{db: s.db}
	}

	return s.authors
}

func (s *SQLiteStore) Genres() storer.GenreRepository {
	if s.genres == nil {
		s.genres = &GenreRepository{db: s.db}
	}

	return s.genres
}

func (s *SQLiteStore) BooksAuthors() storer.BookAuthorRepository {
	if s.booksAuthors == nil {
		s.booksAuthors = &BooksAuthorsRepository{db: s.db}
	}

	return s.booksAuthors
}

func (s *SQLiteStore) Close() error {
	if s == nil {
		return storer.ErrStoreNilOrEmpty
	}

	return s.db.Close()
}

func NewStore(c *config.Config) (storer.Storer, error) {
	if c == nil {
		return nil, config.ErrEmptyConfig
	}

	dburl := fmt.Sprintf("file:%s?_foreign_keys=on", c.DBURL)

	db, err := sqlx.Connect(c.DBType, dburl)
	if err != nil {
		return nil, err
	}

	return &SQLiteStore{
		db: db,
	}, nil
}
