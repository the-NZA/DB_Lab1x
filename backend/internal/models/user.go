package models

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID             string `json:"id" db:"id"`
	Username       string `json:"username" db:"username"`
	HashedPassword string `json:"-" db:"passwd"`
	Email          string `json:"email" db:"email"`
	IsAdmin        bool   `json:"is_admin" db:"is_admin"`
	Deleted        bool   `json:"deleted" db:"deleted"`
}

func (u User) Validate() error {
	return validation.ValidateStruct(&u,
		validation.Field(&u.Username, validation.Required, validation.Length(8, 0)),
		validation.Field(&u.HashedPassword, validation.Required, validation.Length(8, 0)),
		validation.Field(&u.Email, validation.When(u.Email != "", validation.Required, is.Email)),
	)
}

func (u *User) HashPassword(password string) error {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return err
	}

	u.HashedPassword = string(hashedBytes)

	return nil
}

func (u User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.HashedPassword), []byte(password))

	return err == nil
}
