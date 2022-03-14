package models

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type Genre struct {
	ID      string `json:"id" db:"id"`
	Title   string `json:"title" db:"title"`
	Snippet string `json:"snippet" db:"snippet"`
	Deleted bool   `json:"deleted" db:"deleted"`
}

// Validate fields which must always have values
func (g Genre) Validate() error {
	return validation.ValidateStruct(&g,
		validation.Field(&g.ID, validation.When(g.ID != "", validation.Required, validation.Length(1, 0))),
		validation.Field(&g.Title, validation.Required, validation.Length(1, 0)),
		validation.Field(&g.Snippet, validation.When(g.Snippet != "", validation.Required, validation.Length(1, 0))),
		validation.Field(&g.Deleted, validation.In(true, false)),
	)
}
