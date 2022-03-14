package dblab

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/the-NZA/DB_Lab1/backend/internal/models"
)

var (
	ErrNoIDSpecified  = errors.New("You must specify ID in URL param")
	ErrNoQueryParams  = errors.New("You must specify at least 1 query param")
	ErrLenQueryParams = errors.New("Query param length must be 1 or greater")
)

/*
* Books endpoints
 */
// handles GET /api/book/:bookID
func (a *App) handleBookGet() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "bookID")
		if id == "" {
			a.error(w, r, http.StatusInternalServerError, ErrNoIDSpecified)
			return
		}

		// Try get book
		book, err := a.services.BookService().Get(id)
		if err != nil {
			a.logger.Logf("[INFO] During book getting: %v\n", err)
			a.error(w, r, http.StatusInternalServerError, err)
			return
		}

		a.respond(w, r, http.StatusOK, book)
	}
}

// handles POST /api/book
func (a *App) handleBookAdd() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			req models.BookWithAuthors
			err error
		)

		if err = json.NewDecoder(r.Body).Decode(&req); err != nil {
			a.logger.Logf("[INFO] During body parse: %v\n", err)
			a.error(w, r, http.StatusBadRequest, err)
			return
		}

		// Reset book ID
		req.Book.ID = ""

		// Try save new book with authors IDs
		req, err = a.services.BookService().Add(req)
		if err != nil {
			a.logger.Logf("[INFO] During book saving: %v\n", err)
			a.error(w, r, http.StatusInternalServerError, err)
			return
		}

		a.respond(w, r, http.StatusCreated, req)
	}
}

// handles PUT /api/book
func (a *App) handleBookUpdate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			req models.BookWithAuthors
			err error
		)

		if err = json.NewDecoder(r.Body).Decode(&req); err != nil {
			a.logger.Logf("[INFO] During body parse: %v\n", err)
			a.error(w, r, http.StatusBadRequest, err)
			return
		}

		// Try update book
		req, err = a.services.BookService().Update(req)
		if err != nil {
			a.logger.Logf("[INFO] During book updating: %v\n", err)
			a.error(w, r, http.StatusInternalServerError, err)
			return
		}

		a.respond(w, r, http.StatusOK, req)
	}
}

// handles DELETE /api/book/:bookID
func (a *App) handleBookDelete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "bookID")
		if id == "" {
			a.error(w, r, http.StatusInternalServerError, ErrNoIDSpecified)
			return
		}

		// Try delete book by ID
		err := a.services.BookService().Delete(id)
		if err != nil {
			a.logger.Logf("[INFO] During book deleting: %v\n", err)
			a.error(w, r, http.StatusInternalServerError, err)
			return
		}

		a.respond(w, r, http.StatusOK, "Deleted")
	}
}

// handles GET /api/book/all
func (a *App) handleBookGetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Try get all books
		books, err := a.services.BookService().GetAll()
		if err != nil {
			a.logger.Logf("[INFO] During all book getting: %v\n", err)
			a.error(w, r, http.StatusInternalServerError, err)
			return
		}

		a.respond(w, r, http.StatusOK, books)
	}
}

/*
* Genres endpoints
 */
// handles GET /api/genre/:genreID
func (a *App) handleGenreGet() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "genreID")
		if id == "" {
			a.error(w, r, http.StatusInternalServerError, ErrNoIDSpecified)
			return
		}

		// Try get genre
		genre, err := a.services.GenreService().Get(id)
		if err != nil {
			a.logger.Logf("[INFO] During genre getting: %v\n", err)
			a.error(w, r, http.StatusInternalServerError, err)
			return
		}

		a.respond(w, r, http.StatusOK, genre)
	}
}

// handles POST /api/genre
func (a *App) handleGenreAdd() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			genre models.Genre
			err   error
		)

		if err = json.NewDecoder(r.Body).Decode(&genre); err != nil {
			a.logger.Logf("[INFO] During body parse: %v\n", err)
			a.error(w, r, http.StatusBadRequest, err)
			return
		}

		// Try save new genre
		genre, err = a.services.GenreService().Add(genre)
		if err != nil {
			a.logger.Logf("[INFO] During genre saving: %v\n", err)
			a.error(w, r, http.StatusInternalServerError, err)
			return
		}

		a.respond(w, r, http.StatusCreated, genre)
	}
}

// handles PUT /api/genre
func (a *App) handleGenreUpdate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			genre models.Genre
			err   error
		)

		if err = json.NewDecoder(r.Body).Decode(&genre); err != nil {
			a.logger.Logf("[INFO] During body parse: %v\n", err)
			a.error(w, r, http.StatusBadRequest, err)
			return
		}

		// Try update genre
		genre, err = a.services.GenreService().Update(genre)
		if err != nil {
			a.logger.Logf("[INFO] During genre updating: %v\n", err)
			a.error(w, r, http.StatusInternalServerError, err)
			return
		}

		a.respond(w, r, http.StatusOK, genre)
	}
}

// handles DELETE /api/genre/:genreID
func (a *App) handleGenreDelete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "genreID")
		if id == "" {
			a.error(w, r, http.StatusInternalServerError, ErrNoIDSpecified)
			return
		}

		// Try delete genre by ID
		err := a.services.GenreService().Delete(id)
		if err != nil {
			a.logger.Logf("[INFO] During genre deleting: %v\n", err)
			a.error(w, r, http.StatusInternalServerError, err)
			return
		}

		a.respond(w, r, http.StatusOK, "Deleted")
	}
}

// handles GET /api/genre/all
func (a *App) handleGenreGetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Try get all genres
		genres, err := a.services.GenreService().GetAll()
		if err != nil {
			a.logger.Logf("[INFO] During all genre getting: %v\n", err)
			a.error(w, r, http.StatusInternalServerError, err)
			return
		}

		a.respond(w, r, http.StatusOK, genres)
	}
}

/*
* Authors endpoints
 */
// handles GET /api/author/:authorID
func (a *App) handleAuthorGet() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "authorID")
		if id == "" {
			a.error(w, r, http.StatusInternalServerError, ErrNoIDSpecified)
			return
		}

		// Try get author
		author, err := a.services.AuthorService().Get(id)
		if err != nil {
			a.logger.Logf("[INFO] During author getting: %v\n", err)
			a.error(w, r, http.StatusInternalServerError, err)
			return
		}

		a.respond(w, r, http.StatusOK, author)
	}
}

// handles POST /api/author
func (a *App) handleAuthorAdd() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			req models.AuthorWithBooks
			err error
		)

		if err = json.NewDecoder(r.Body).Decode(&req); err != nil {
			a.logger.Logf("[INFO] During body parse: %v\n", err)
			a.error(w, r, http.StatusBadRequest, err)
			return
		}

		// Reset ID and Deleted flag for new author
		req.Author.ID, req.Author.Deleted = "", false

		// Try save new author
		req, err = a.services.AuthorService().Add(req)
		if err != nil {
			a.logger.Logf("[INFO] During author saving: %v\n", err)
			a.error(w, r, http.StatusInternalServerError, err)
			return
		}

		a.respond(w, r, http.StatusCreated, req)
	}
}

// handles PUT /api/author
func (a *App) handleAuthorUpdate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			req models.AuthorWithBooks
			err error
		)

		if err = json.NewDecoder(r.Body).Decode(&req); err != nil {
			a.logger.Logf("[INFO] During body parse: %v\n", err)
			a.error(w, r, http.StatusBadRequest, err)
			return
		}

		// Try update author
		req, err = a.services.AuthorService().Update(req)
		if err != nil {
			a.logger.Logf("[INFO] During author updating: %v\n", err)
			a.error(w, r, http.StatusInternalServerError, err)
			return
		}

		a.respond(w, r, http.StatusOK, req)
	}
}

// handles DELETE /api/author/:authorID
func (a *App) handleAuthorDelete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "authorID")
		if id == "" {
			a.error(w, r, http.StatusInternalServerError, ErrNoIDSpecified)
			return
		}

		// Try delete author by id
		err := a.services.AuthorService().Delete(id)
		if err != nil {
			a.logger.Logf("[INFO] During author deleting: %v\n", err)
			a.error(w, r, http.StatusInternalServerError, err)
			return
		}

		a.respond(w, r, http.StatusOK, "Deleted")
	}
}

// handles GET /api/author/all
func (a *App) handleAuthorGetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Try get all authors
		authors, err := a.services.AuthorService().GetAll()
		if err != nil {
			a.logger.Logf("[INFO] During all author getting: %v\n", err)
			a.error(w, r, http.StatusInternalServerError, err)
			return
		}

		a.respond(w, r, http.StatusOK, authors)
	}
}

/*
* BookAuthor handlers
 */

// handle /api/book-author?author_id=AuthorID&book_id=BookID
func (a *App) handleGetBookAuthor() http.HandlerFunc {
	const (
		authorParam = "author_id"
		bookParam   = "book_id"
	)

	return func(w http.ResponseWriter, r *http.Request) {
		// hasAuthorID := r.URL.Query().Has(authorParam)
		// hasBookID := r.URL.Query().Has(bookParam)

		// if !hasAuthorID && !hasBookID {
		// 	a.error(w, r, http.StatusBadRequest, ErrNoQueryParams)
		// 	return
		// }

		// if hasAuthorID && len(authorID) == 0 || hasBookID && len(bookID) == 0 {
		// 	a.error(w, r, http.StatusBadRequest, ErrLenQueryParams)
		// 	return
		// }

		authorID := r.URL.Query().Get(authorParam)
		bookID := r.URL.Query().Get(bookParam)

		booksAuthors, err := a.services.BooksAuthors().GetByIDs(bookID, authorID)
		if err != nil {
			a.error(w, r, http.StatusInternalServerError, err)
			return
		}

		a.respond(w, r, http.StatusOK, booksAuthors)
	}
}
