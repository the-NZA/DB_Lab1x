package models

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type Book struct {
	ID          string `json:"id" db:"id"`
	Title       string `json:"title" db:"title"`
	Snippet     string `json:"snippet" db:"snippet"`
	GenreID     string `json:"genre_id" db:"genre_id"`
	PagesCnt    uint   `json:"pages_cnt" db:"pages_cnt"`
	PublishYear int    `json:"pub_year" db:"pub_year"`
	Deleted     bool   `json:"deleted" db:"deleted"`
}

// Validate fields which must always have values
func (b Book) Validate() error {
	return validation.ValidateStruct(&b,
		validation.Field(&b.ID, validation.When(b.ID != "", validation.Required, validation.Length(1, 0))),
		validation.Field(&b.Snippet, validation.When(b.Snippet != "", validation.Required, validation.Length(1, 0))),
		validation.Field(&b.Title, validation.Required, validation.Length(1, 0)),
		validation.Field(&b.GenreID, validation.Required),
		validation.Field(&b.PagesCnt, validation.When(b.PagesCnt != 0, validation.Required)),
		validation.Field(&b.PublishYear, validation.Required),
		validation.Field(&b.Deleted, validation.In(true, false)),
	)
}

// BookWithAuthors used for creating or updating books
type BookWithAuthors struct {
	Book       Book     `json:"book"`
	AuthorsIDs []string `json:"authors_ids"`
}

func (bwa BookWithAuthors) Validate() error {
	err := bwa.Book.Validate()
	if err != nil {
		return err
	}

	return validation.ValidateStruct(&bwa,
		validation.Field(&bwa.AuthorsIDs, validation.Required, validation.Length(1, 0)),
	)
}
