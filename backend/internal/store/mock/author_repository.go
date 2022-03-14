package mock

import (
	"fmt"
	"math/rand"
	"strconv"

	"github.com/the-NZA/DB_Lab1/backend/internal/models"
)

var authors = []models.Author{
	{
		ID:        "1",
		Firstname: "Ivan",
		Lastname:  "Ivanov",
		Surname:   "Ivanovich",
		Snippet:   "This is one of the best Ivan Ivanov in the whole world",
		Deleted:   false,
	},
	{
		ID:        "2",
		Firstname: "Semen",
		Lastname:  "Semenov",
		Surname:   "Semenovich",
		Snippet:   "This is one of the best Semen Semenov in the whole world",
		Deleted:   false,
	},
	{
		ID:        "3",
		Firstname: "Kirill",
		Lastname:  "Kirillov",
		Surname:   "Kirillovich",
		Snippet:   "This is one of the best Kirill Kirillov in the whole world",
		Deleted:   false,
	},
}

// Get next ID
func getNextAuthorID(a []models.Author) int {
	if len(a) < 1 {
		return 0
	}

	next, err := strconv.Atoi(a[0].ID)
	if err != nil {
		return (rand.Int() % len(a)) + 1
	}

	for i := range a {
		n, err := strconv.Atoi(a[i].ID)
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

type AuthorRepository struct{}

// find one author by passed ID
func (a AuthorRepository) findByID(ID string) (int, error) {
	for i := range authors {
		if authors[i].ID == ID {
			return i, nil
		}
	}

	return -1, fmt.Errorf("Author not found")
}

// Get one author from authors
func (a AuthorRepository) Get(ID string) (models.Author, error) {
	idx, err := a.findByID(ID)
	if err != nil {
		return models.Author{}, fmt.Errorf("Author not found")
	}

	return authors[idx], nil
}

// Add one author to authors
func (a AuthorRepository) Add(author models.AuthorWithBooks) (models.AuthorWithBooks, error) {
	newID := getNextAuthorID(authors)

	author.Author.ID = strconv.Itoa(newID)

	authors = append(authors, author.Author)

	return author, nil
}

// Update one author in authors
func (a AuthorRepository) Update(author models.AuthorWithBooks) (models.AuthorWithBooks, error) {
	idx, err := a.findByID(author.Author.ID)
	if err != nil {
		return author, fmt.Errorf("Author not found")
	}

	authors[idx] = author.Author

	return author, nil
}

// Delete one author from authors
func (a AuthorRepository) Delete(ID string) error {
	idx, err := a.findByID(ID)
	if err != nil {
		return fmt.Errorf("Author not found")
	}

	authors[idx] = authors[len(authors)-1]
	authors = authors[:len(authors)-1]

	return nil
}

// Get all authors from authors
func (a AuthorRepository) GetAll() ([]models.Author, error) {
	if len(authors) < 1 {
		return nil, fmt.Errorf("Authors not found")
	}

	return authors, nil
}
