package app

import (
	"context"
	"net/http"
	"time"

	"01.alem.school/git/abdu0222/forum/pkg"
)

var (
	WelcomeCookieOnPaths = make(map[string]struct{})
	HomeCookieOnPaths    = make(map[string]struct{})
)

func AddWelcomeCookieCheckOnPaths(paths ...string) {
	for _, path := range paths {
		WelcomeCookieOnPaths[path] = struct{}{}
	}
}

func AddHomeCookieCheckOnPaths(paths ...string) {
	for _, path := range paths {
		HomeCookieOnPaths[path] = struct{}{}
	}
}

const Username string = ""

func (app *App) WelcomeMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if _, ok := WelcomeCookieOnPaths[r.URL.Path]; !ok {
			pkg.ErrorHandler(w, http.StatusNotFound)
			return
		} else {
			c, err := r.Cookie("session_token")
			if err == http.ErrNoCookie {
				next.ServeHTTP(w, r)
				return
			}
			sessionFromDb, err := app.sessionService.GetSessionByToken(c.Value)
			if err != nil {
				next.ServeHTTP(w, r)
				return
			}
			if sessionFromDb.Expiry.Before(time.Now()) {
				next.ServeHTTP(w, r)
				return
			}
			http.Redirect(w, r, "/", http.StatusFound)
			return
		}
	}
}

func (app *App) HomeMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if _, ok := HomeCookieOnPaths[r.URL.Path]; !ok {
			pkg.ErrorHandler(w, http.StatusNotFound)
			return
		} else {
			c, err := r.Cookie("session_token")
			if err == http.ErrNoCookie {
				http.Redirect(w, r, "/sign-in-form", http.StatusFound)
				return
			}
			sessionFromDb, err := app.sessionService.GetSessionByToken(c.Value)
			if err != nil {
				http.Redirect(w, r, "/sign-in-form", http.StatusFound)
				return
			}
			if sessionFromDb.Expiry.Before(time.Now()) {
				http.Redirect(w, r, "/sign-in-form", http.StatusFound)
				return
			}
			next.ServeHTTP(w, r)
		}
	}
}

func (app *App) authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		c, err := r.Cookie("session_token")
		if err != nil {
			http.Redirect(w, r, "sign-in-form", http.StatusFound)
			return
		}
		token := c.Value
		sessionFromDb, err := app.sessionService.GetSessionByToken(token)
		if err != nil {
			http.Redirect(w, r, "sign-in-form", http.StatusFound)
			return
		}
		if sessionFromDb.Expiry.Before(time.Now()) {
			http.Redirect(w, r, "/sign-in-form", http.StatusFound)
			return
		}
		ctx := context.WithValue(r.Context(), Username, sessionFromDb.Username)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
