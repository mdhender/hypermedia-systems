// Copyright (c) 2023 Michael D Henderson. All rights reserved.

package contacts

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

func (a *App) getContacts() http.HandlerFunc {
	t, err := template.ParseFiles(filepath.Join(a.templates, "contacts.gohtml"))
	if err != nil {
		log.Printf("[contacts] getContacts: %v\n", err)
		return func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
	}
	type payload struct {
		Search   string
		Contacts []Contact
	}
	// log.Printf("[contacts] getContacts %+v\n", a.contacts.contacts)

	return func(w http.ResponseWriter, r *http.Request) {
		// a.contacts.Dump(os.Stdout)
		var p payload
		p.Search = r.URL.Query().Get("q")
		// log.Printf("[contacts] getContacts: pSearch %+v\n", pSearch)
		bb := &bytes.Buffer{}
		var results *Contacts
		if p.Search == "" { // missing or explicitly zero length
			results = a.contacts.All()
		} else {
			// log.Printf("[contacts] getContacts: search %+v\n", search)
			results = a.contacts.Search(p.Search)
		}
		for _, contact := range results.contacts {
			p.Contacts = append(p.Contacts, *contact)
		}
		if err = t.Execute(bb, p); err != nil {
			log.Printf("%s %s: %v\n", r.Method, r.URL, err)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		_, _ = w.Write(bb.Bytes())
	}
}

func (a *App) render_template(name string, data any) ([]byte, error) {
	return []byte("Hello world"), nil
}
