package main

import (
	"math/rand/v2"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
)

const (
	maxMilliseconds = 2000
)

func delayWidgets(ms int) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			time.Sleep(time.Duration(rand.IntN(ms)) * time.Millisecond)

			next.ServeHTTP(w, r)
		})
	}
}

func (app *Application) setRoutes(mux *chi.Mux) {
	// Pages.
	mux.Get("/", app.root)

	// Widgets.

	// Widgets that we delay.
	mux.Group(func(mux chi.Router) {
		mux.Use(delayWidgets(maxMilliseconds))
	})
}
