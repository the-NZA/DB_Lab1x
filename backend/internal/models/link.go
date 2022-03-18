package models

import validation "github.com/go-ozzo/ozzo-validation/v4"

type Link struct {
	ID      string `json:"id" db:"id"`
	Link    string `json:"link" db:"link"`
	Deleted bool   `json:"deleted" db:"deleted"`
	BookID  string `json:"book_id" db:"book_id"`
}

// Validate fields which must always have values
func (l Link) Validate() error {
	return validation.ValidateStruct(&l,
		validation.Field(&l.Link, validation.Length(5, 0)),
		validation.Field(&l.Deleted, validation.In(true, false)),
	)
}
