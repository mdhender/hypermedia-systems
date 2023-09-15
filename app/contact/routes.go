// Copyright (c) 2023 Michael D Henderson. All rights reserved.

// Package contact implements the Contact.app server from the Hypermedia Systems book.
package contact

import (
	"bytes"
	"github.com/matryer/way"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

func (a *App) Router() http.Handler {
	router := way.NewRouter()
	router.Handle("GET", "/", a.getIndex())
	router.Handle("GET", "/contact", a.getContacts())
	return router
}

func (a *App) getContacts() http.HandlerFunc {
	t, err := template.ParseFiles(filepath.Join(a.templates, "contact.gohtml"))
	if err != nil {
		log.Printf("get contact: %v\n", err)
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		})
	}
	type payload struct {
		Search   string
		Contacts []string
	}
	// log.Printf("[contact] contact %+v\n", a.contact.contact)

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// a.contact.Dump(os.Stdout)
		params := r.URL.Query()
		// log.Printf("get contact: params %+v\n", params)
		bb := &bytes.Buffer{}
		var p payload
		if search, ok := params["q"]; !ok {
			c := a.contacts.All()
			for _, contact := range c.contacts {
				p.Contacts = append(p.Contacts, contact.Name)
			}
		} else if len(search) != 1 {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		} else if len(search[0]) == 0 {
			// log.Printf("[contact] get contact: search %+v\n", search)
			c := a.contacts.All()
			for _, contact := range c.contacts {
				p.Contacts = append(p.Contacts, contact.Name)
			}
		} else {
			// log.Printf("[contact] get contact: search %+v\n", search)
			p.Search = search[0]
			c := a.contacts.Search(search[0])
			for _, contact := range c.contacts {
				p.Contacts = append(p.Contacts, contact.Name)
			}
		}
		if err = t.Execute(bb, p); err != nil {
			log.Printf("%s %s: %v\n", r.Method, r.URL, err)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		_, _ = w.Write(bb.Bytes())
	})
}

func (a *App) getIndex() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/contact", http.StatusTemporaryRedirect)
	})
}

func (a *App) render_template(name string, data any) ([]byte, error) {
	return []byte("Hello world"), nil
}
