package services

import (
	"errors"

	"github.com/the-NZA/DB_Lab1x/backend/internal/models"
	"github.com/the-NZA/DB_Lab1x/backend/internal/store/storer"
)

var (
	ErrWrongUserCredentials = errors.New("Something wrong with login or password")
)

type AuthService struct {
	repository storer.UserRepository
}

func (a *AuthService) Login(username string, password string) (models.User, error) {
	user, err := a.repository.GetByUsername(username)
	if err != nil {
		return user, err
	}

	if !user.CheckPassword(password) {
		return models.User{}, ErrWrongUserCredentials
	}

	return user, nil
}

func (a *AuthService) Register(user models.User) (models.User, error) {
	if err := user.Validate(); err != nil {
		return user, err
	}

	return a.repository.Add(user)
}
