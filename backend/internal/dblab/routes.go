package dblab

import (
	"net/http"
)

func (a *App) handleIndexPage() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		a.respond(w, r, http.StatusOK, "This is index page")
	}
}
