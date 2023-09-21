// Copyright (c) 2023 Michael D Henderson. All rights reserved.

package web1

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/mdhender/hypermedia-systems/internal/config"
	"github.com/mdhender/hypermedia-systems/internal/mw"
	"net/http"
)

func (a *App) Router(cfg *config.Config) http.Handler {
	r := chi.NewRouter()

	// add global middleware
	if cfg.Middleware.BadRunes {
		r.Use(mw.BadRunes)
	}
	if cfg.Middleware.Cors {
		r.Use(mw.CORS)
	}
	if cfg.Middleware.Logging {
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
