package storer

import "github.com/the-NZA/DB_Lab1x/backend/internal/models"

type BookReporsitory interface {
	Get(string) (models.Book, error)
	Add(models.BookWithAuthors) (models.BookWithAuthors, error)
	Update(models.BookWithAuthors) (models.BookWithAuthors, error)
	Delete(string) error
	GetAll() ([]models.Book, error)
	GetRandom3() ([]models.Book, error)
	Search(title, author, genre string) ([]models.Book, error)
}

type AuthorRepository interface {
	Get(string) (models.Author, error)
	Add(models.AuthorWithBooks) (models.AuthorWithBooks, error)
	Update(models.AuthorWithBooks) (models.AuthorWithBooks, error)
	Delete(string) error
	GetAll() ([]models.Author, error)
}

type GenreRepository interface {
	Get(string) (models.Genre, error)
	Add(models.Genre) (models.Genre, error)
	Update(models.Genre) (models.Genre, error)
	Delete(string) error
	GetAll() ([]models.Genre, error)
}

type BookAuthorRepository interface {
	GetByIDs(string, string) ([]models.BookAuthor, error)
}

type UserRepository interface {
	GetByUsername(string) (models.User, error)
	Add(models.User) (models.User, error)
}

type LinkRepository interface {
	GetByBookID(bookID string) ([]models.Link, error)
	Add(link models.Link) (models.Link, error)
	Delete(linkID string) error
}
