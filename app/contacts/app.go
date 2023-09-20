// Copyright (c) 2023 Michael D Henderson. All rights reserved.

// Package contacts implements the Contacts.app server from the Hypermedia Systems book.
package contacts

import (
	"github.com/mdhender/hypermedia-systems/internal/store"
	"log"
	"net"
	"net/http"
	"time"
)

type App struct {
	router    http.Handler
	templates string
	server    struct {
		http.Server
		host       string
		port       string
		middleware struct {
			badRunes func(http.Handler) http.Handler
			cors     func(http.Handler) http.Handler
			logging  func(http.Handler) http.Handler
		}
	}
	store *store.Store
}

func New(options ...Option) (*App, error) {
	a := &App{
		templates: ".",
	}
	a.server.host = ""
	a.server.port = "8080"
	a.server.Addr = net.JoinHostPort(a.server.host, a.server.port)
	a.server.MaxHeaderBytes = 1 << 20 // 1mb?
	a.server.ReadTimeout = 5 * time.Second
	a.server.WriteTimeout = 10 * time.Second

	for _, option := range options {
		if err := option(a); err != nil {
			return nil, err
		}
	}

	if a.store == nil {
		a.store = store.New()
	}

	a.router = a.Router()

	return a, nil
}

func (a *App) ListenAndServe() error {
	log.Printf("[app] listening on %q\n", a.server.Addr)
	return a.server.ListenAndServe()
}
