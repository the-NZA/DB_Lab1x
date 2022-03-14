package storer

import "errors"

var (
	ErrStoreNilOrEmpty = errors.New("Store must be initialized before usage")
)

type Storer interface {
	Books() BookReporsitory
	Authors() AuthorRepository
	Genres() GenreRepository
	BooksAuthors() BookAuthorRepository
	Close() error
}
