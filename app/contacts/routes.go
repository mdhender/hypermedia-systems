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
	r.Get("/contacts", a.contacts())
	r.Get("/contacts/new", a.contacts_new_get())
	r.Post("/contacts/new", a.contacts_new())
	r.Get("/contacts/{id}", a.contacts_view())
	r.Get("/contacts/{id}/edit", a.contacts_edit_get())
	r.Post("/contacts/{id}/edit", a.contacts_edit_post())
	r.Post("/contacts/{id}/delete", a.contacts_delete())

	a.server.Handler = r
	return a.server.Handler
}
