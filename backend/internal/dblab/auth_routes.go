package dblab

import (
	"encoding/json"
	"net/http"
)

// handles POST /auth/login
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

		a.respond(w, r, http.StatusCreated, user)
		// a.respond(w, r, http.StatusOK, "Login successful")
	}
}

// handles POST /auth/signup
func (a *App) handleRegister() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// var (
		// 	req models.BookWithAuthors
		// 	err error
		// )

		// if err = json.NewDecoder(r.Body).Decode(&req); err != nil {
		// 	a.logger.Logf("[INFO] During body parse: %v\n", err)
		// 	a.error(w, r, http.StatusBadRequest, err)
		// 	return
		// }

		// // Reset book ID
		// req.Book.ID = ""

		// // Try save new book with authors IDs
		// req, err = a.services.BookService().Add(req)
		// if err != nil {
		// 	a.logger.Logf("[INFO] During book saving: %v\n", err)
		// 	a.error(w, r, http.StatusInternalServerError, err)
		// 	return
		// }

		// a.respond(w, r, http.StatusCreated, req)
		a.respond(w, r, http.StatusOK, "Register")
	}
}
