package mock

import (
	"fmt"
	"math/rand"
	"strconv"

	"github.com/the-NZA/DB_Lab1/backend/internal/models"
)

var books = []models.Book{
	{
		ID:          "1",
		Title:       "First",
		Snippet:     "First book description",
		GenreID:     "1",
		PagesCnt:    100,
		PublishYear: 2013,
		Deleted:     false,
	},
	{
		ID:          "2",
		Title:       "Second",
		Snippet:     "Second book description",
		GenreID:     "2",
		PagesCnt:    120,
		PublishYear: 1993,
		Deleted:     false,
	},
	{
		ID:          "3",
		Title:       "Third",
		Snippet:     "Third book description",
		GenreID:     "1",
		PagesCnt:    110,
		PublishYear: 2015,
		Deleted:     false,
	},
}

// Get next ID
func getNextBookID(b []models.Book) int {
	if len(b) < 1 {
		return 0
	}

	next, err := strconv.Atoi(b[0].ID)
	if err != nil {
		return (rand.Int() % len(b)) + 1
	}

	for i := range b {
		n, err := strconv.Atoi(b[i].ID)
		if err != nil {
			continue
		}

		if n > next {
			next = n
		}
	}

	next++
	return next
}

type BookRepository struct{}

// find one book by passed ID
func (b BookRepository) findByID(ID string) (int, error) {
	for i := range books {
		if books[i].ID == ID {
			return i, nil
		}
	}

	return -1, fmt.Errorf("Book not found")
}

// Get one book from books
func (b BookRepository) Get(ID string) (models.Book, error) {
	idx, err := b.findByID(ID)
	if err != nil {
		return models.Book{}, fmt.Errorf("Book not found")
	}

	return books[idx], nil
}

// Add one book to books
func (b BookRepository) Add(ba models.BookWithAuthors) (models.BookWithAuthors, error) {
	newID := getNextBookID(books)

	ba.Book.ID = strconv.Itoa(newID)

	books = append(books, ba.Book)

	return ba, nil
}

// Update one book in books
func (b BookRepository) Update(ba models.BookWithAuthors) (models.BookWithAuthors, error) {
	idx, err := b.findByID(ba.Book.ID)
	if err != nil {
		return ba, fmt.Errorf("Book not found")
	}

	books[idx] = ba.Book

	return ba, nil
}

// Delete one book from books
func (b BookRepository) Delete(ID string) error {
	idx, err := b.findByID(ID)
	if err != nil {
		return fmt.Errorf("Book not found")
	}

	books[idx] = books[len(books)-1]
	books = books[:len(books)-1]

	return nil
}

// Gell all books from books
func (b BookRepository) GetAll() ([]models.Book, error) {
	if len(books) < 1 {
		return nil, fmt.Errorf("Books not found")
	}

	return books, nil
}
