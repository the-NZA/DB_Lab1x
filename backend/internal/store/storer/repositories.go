package storer

import "github.com/the-NZA/DB_Lab1/backend/internal/models"

type BookReporsitory interface {
	Get(string) (models.Book, error)
	Add(models.BookWithAuthors) (models.BookWithAuthors, error)
	Update(models.BookWithAuthors) (models.BookWithAuthors, error)
	Delete(string) error
	GetAll() ([]models.Book, error)
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
