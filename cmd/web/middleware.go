package main

import (
	"github.com/justinas/nosurf"
	"net/http"
)

// NoSurf Adds CSRF protection to all post requests
func NoSurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)
	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   app.InProduction,
		SameSite: http.SameSiteLaxMode,
	})
	return csrfHandler
}

// SessionLoad Loads and Saves Sessions Every Request
func SessionLoad(next http.Handler) http.Handler {
	return session.LoadAndSave(next)
}
