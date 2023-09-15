// Copyright (c) 2023 Michael D Henderson. All rights reserved.

package contacts

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"net/http"
)

func (a *App) Router() http.Handler {
	r := chi.NewRouter()

	// add global middleware
	if a.server.middleware.badRunes != nil {
		r.Use(a.server.middleware.badRunes)
	}
	if a.server.middleware.cors != nil {
		r.Use(a.server.middleware.cors)
	}
	if a.server.middleware.logging != nil {
		r.Use(middleware.Logger)
	}

	// create routes
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/contacts", http.StatusTemporaryRedirect)
	})
	r.Get("/contacts", a.getContacts())

	a.server.Handler = r
	return a.server.Handler
}
