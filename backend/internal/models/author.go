package models

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type Author struct {
	ID        string `json:"id" db:"id"`
	Firstname string `json:"firstname" db:"firstname"`
	Lastname  string `json:"lastname" db:"lastname"`
	Surname   string `json:"surname" db:"surname"`
	Snippet   string `json:"snippet" db:"snippet"`
	Deleted   bool   `json:"deleted" db:"deleted"`
}

// Validate fields which must always have values
func (a Author) Validate() error {
	return validation.ValidateStruct(&a,
		validation.Field(&a.Firstname, validation.Required, validation.Length(1, 0)),
		validation.Field(&a.Lastname, validation.Required, validation.Length(1, 0)),
		validation.Field(&a.Surname, validation.When(a.Surname != "", validation.Required, validation.Length(1, 0))),
		validation.Field(&a.Snippet, validation.When(a.Snippet != "", validation.Required, validation.Length(1, 0))),
		validation.Field(&a.Deleted, validation.In(true, false)),
	)
}

type AuthorWithBooks struct {
	Author   Author   `json:"author"`
	BooksIDs []string `json:"books_ids"`
}

func (awb AuthorWithBooks) Validate() error {
	err := awb.Author.Validate()
	if err != nil {
		return err
	}

	return validation.ValidateStruct(&awb,
		validation.Field(&awb.BooksIDs, validation.NotNil))
}
