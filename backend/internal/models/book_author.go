package models

import validation "github.com/go-ozzo/ozzo-validation/v4"

type BookAuthor struct {
	ID       string `json:"id" db:"id"`
	BookID   string `json:"book_id" db:"book_id"`
	AuthorID string `json:"author_id" db:"author_id"`
	Deleted  bool   `json:"deleted" db:"deleted"`
}

// Validate fields of BookAuthor struct
func (ba BookAuthor) Validate() error {
	return validation.ValidateStruct(&ba,
		validation.Field(&ba.ID, validation.When(ba.ID != "", validation.Required, validation.Length(1, 0))),
		validation.Field(&ba.BookID, validation.Required, validation.Length(1, 0)),
		validation.Field(&ba.AuthorID, validation.Required, validation.Length(1, 0)),
	)
}
