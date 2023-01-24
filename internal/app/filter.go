package app

import (
	"net/http"

	"github.com/with-insomnia/Forum-Golang/pkg"
)

func (app *App) FilterHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		pkg.ErrorHandler(w, http.StatusMethodNotAllowed)
		return
	}

	switch r.URL.Path {
	case "/post/filter":
		// for auth users
	case "/welcome/post/filter":
		// for unauth users
	default:
		pkg.ErrorHandler(w, http.StatusNotFound)
		return
	}
}

func (app *App) FilterWelcomeHandler(w http.ResponseWriter, r *http.Request) {
}
