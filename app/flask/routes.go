// Copyright (c) 2023 Michael D Henderson. All rights reserved.

// Package flask implements a sample application.
package flask

import (
	"fmt"
	"github.com/matryer/way"
	"net/http"
)

func (a *App) Router() http.Handler {
	router := way.NewRouter()
	router.Handle("GET", "/", a.getIndex())
	router.Handle("GET", "/contacts", a.getContacts())
	return router
}

func (a *App) getIndex() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/contacts", http.StatusTemporaryRedirect)
	})
}

func (a *App) getContacts() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, _ = fmt.Fprintf(w, "hello")
	})
}
