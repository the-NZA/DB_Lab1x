package dblab

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/the-NZA/DB_Lab1x/backend/internal/models"
)

// handles POST /auth/login.
func (a *App) handleLogin() http.HandlerFunc {
	type reqLogin struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		var (
			req reqLogin
			err error
		)

		if err = json.NewDecoder(r.Body).Decode(&req); err != nil {
			a.logger.Logf("[INFO] During body parse: %v\n", err)
			a.error(w, r, http.StatusBadRequest, err)
			return
		}

		// Try login user
		user, err := a.services.AuthService().Login(req.Username, req.Password)
		if err != nil {
			a.logger.Logf("[INFO] During user login: %v\n", err)
			a.error(w, r, http.StatusInternalServerError, err)
			return
		}

		a.respond(w, r, http.StatusOK, user)
	}
}

// handles POST /auth/signup.
func (a *App) handleRegister() http.HandlerFunc {
	type reqRegister struct {
		Username string `json:"username"`
		Password string `json:"password"`
		Email    string `json:"email"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		var (
			req  reqRegister
			user models.User
			err  error
		)

		if err = json.NewDecoder(r.Body).Decode(&req); err != nil {
			a.logger.Logf("[INFO] During body parse: %v\n", err)
			a.error(w, r, http.StatusBadRequest, err)
			return
		}

		// lazy check password for emptiness
		if len(strings.TrimSpace(req.Password)) < 8 {
			err2 := fmt.Errorf("password must have at least 8 non-white characters")
			a.logger.Logf("[INFO] Check password length: %v\n", err2)
			a.error(w, r, http.StatusBadRequest, err2)
			return
		}

		user.Username = req.Username
		user.Email = req.Email
		err = user.HashPassword(req.Password)
		if err != nil {
			a.logger.Logf("[INFO] During password hashing: %v\n", err)
			a.error(w, r, http.StatusInternalServerError, err)
			return
		}

		// Try to create new user
		user, err = a.services.AuthService().Register(user)
		if err != nil {
			a.logger.Logf("[INFO] During user creating: %v\n", err)
			a.error(w, r, http.StatusInternalServerError, err)
			return
		}

		a.respond(w, r, http.StatusCreated, user)
	}
}
