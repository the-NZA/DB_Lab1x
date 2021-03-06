package services

import "github.com/the-NZA/DB_Lab1x/backend/internal/models"

type Servicer interface {
	BookService() BookServicer
	GenreService() GenreServicer
	AuthorService() AuthorServicer
	AuthService() AuthServicer
	LinksService() LinkServicer
	BooksAuthors() BooksAuthorsServicer
	Close() error
}

type BookServicer interface {
	Get(string) (models.Book, error)
	Add(models.BookWithAuthors) (models.BookWithAuthors, error)
	Update(models.BookWithAuthors) (models.BookWithAuthors, error)
	Delete(string) error
	GetRandom3() ([]models.Book, error)
	GetAll() ([]models.Book, error)
	Search(title, author, genre string) ([]models.Book, error)
}

type AuthorServicer interface {
	Get(string) (models.Author, error)
	Add(models.AuthorWithBooks) (models.AuthorWithBooks, error)
	Update(models.AuthorWithBooks) (models.AuthorWithBooks, error)
	Delete(string) error
	GetAll() ([]models.Author, error)
}

type GenreServicer interface {
	Get(string) (models.Genre, error)
	Add(models.Genre) (models.Genre, error)
	Update(models.Genre) (models.Genre, error)
	Delete(string) error
	GetAll() ([]models.Genre, error)
}

type BooksAuthorsServicer interface {
	GetByIDs(string, string) ([]models.BookAuthor, error)
}

type AuthServicer interface {
	Login(username, password string) (models.User, error)
	Register(models.User) (models.User, error)
}

type LinkServicer interface {
	Get(string) ([]models.Link, error)
	Add(models.Link) (models.Link, error)
	Delete(string) error
}
