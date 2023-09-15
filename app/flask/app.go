// Copyright (c) 2023 Michael D Henderson. All rights reserved.

package flask

import (
	"log"
	"net/http"
	"time"
)

type App struct {
	router    http.Handler
	contacts  *Contacts
	templates string
}

func New(options ...Option) (*App, error) {
	a := &App{
		templates: ".",
	}
	for _, option := range options {
		if err := option(a); err != nil {
			return nil, err
		}
	}
	a.router = a.Router()
	return a, nil
}

func (a *App) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	started := time.Now()
	defer func(method, route string) {
		log.Printf("%s %s: %v\n", method, route, time.Now().Sub(started))
	}(r.Method, r.URL.Path)
	a.router.ServeHTTP(w, r)
}