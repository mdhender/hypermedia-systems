// Copyright (c) 2023 Michael D Henderson. All rights reserved.

// Package contacts implements a sample application.
package contacts

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
	router.Handle("GET", "/contacts", a.getContacts())
	return router
}

func (a *App) getContacts() http.HandlerFunc {
	t, err := template.ParseFiles(filepath.Join(a.templates, "contacts.gohtml"))
	if err != nil {
		log.Printf("get contacts: %v\n", err)
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		})
	}
	type payload struct {
		Search   string
		Contacts []string
	}
	// log.Printf("[contacts] contacts %+v\n", a.contacts.contacts)

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// a.contacts.Dump(os.Stdout)
		params := r.URL.Query()
		// log.Printf("get contacts: params %+v\n", params)
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
			// log.Printf("[contacts] get contacts: search %+v\n", search)
			c := a.contacts.All()
			for _, contact := range c.contacts {
				p.Contacts = append(p.Contacts, contact.Name)
			}
		} else {
			// log.Printf("[contacts] get contacts: search %+v\n", search)
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
		http.Redirect(w, r, "/contacts", http.StatusTemporaryRedirect)
	})
}

func (a *App) render_template(name string, data any) ([]byte, error) {
	return []byte("Hello world"), nil
}
