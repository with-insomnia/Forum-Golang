package delivery

import (
	"net/http"

	"github.com/with-insomnia/Forum-Golang/pkg"
)

func CommentHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		// TODO get
	case http.MethodPost:
		// TODO post
	default:
		pkg.ErrorHandler(w, http.StatusMethodNotAllowed)
		return
	}
}