// Copyright (c) 2023 Michael D Henderson. All rights reserved.

package contacts

import (
	"fmt"
	"github.com/go-chi/chi"
	"github.com/mdhender/hypermedia-systems/internal/store"
	"log"
	"net/http"
	"strconv"
)

type PayloadContact struct {
	Id     int
	First  string
	Last   string
	Phone  string
	Email  string
	Errors struct {
		Record error
		First  error
		Last   error
		Phone  error
		Email  error
	}
}

func (a *App) contacts() http.HandlerFunc {
	type payload struct {
		Search   string
		Contacts []PayloadContact
		Flash    string
	}
	// log.Printf("[contacts] contacts %+v\n", a.contacts.contacts)

	return func(w http.ResponseWriter, r *http.Request) {
		// a.contacts.Dump(os.Stdout)
		var p payload

		if flash, err := r.Cookie("flash"); err == nil {
			http.SetCookie(w, &http.Cookie{
				Path:     "/",
				Name:     "flash",
				MaxAge:   -1,
				HttpOnly: true,
				Secure:   true,
			})
			p.Flash = flash.Value
		}

		p.Search = r.URL.Query().Get("q")
		// log.Printf("[contacts] contacts: pSearch %+v\n", pSearch)
		var results []store.Contact
		if p.Search == "" { // missing or explicitly zero length
			results = a.store.AllContacts()
		} else {
			// log.Printf("[contacts] contacts: search %+v\n", search)
			results = a.store.SearchContacts(p.Search)
		}
		for _, contact := range results {
			p.Contacts = append(p.Contacts, PayloadContact{
				Id:    contact.Id,
				First: contact.First,
				Last:  contact.Last,
				Phone: contact.Phone,
				Email: contact.Email,
			})
		}
		a.render(w, r, p, "layout", "contacts")
	}
}

func (a *App) contacts_new_get() http.HandlerFunc {
	type payload struct {
		Contact PayloadContact
	}
	// log.Printf("[contacts] contacts_new_get %+v\n", a.contacts.contacts)

	return func(w http.ResponseWriter, r *http.Request) {
		p := payload{}
		a.render(w, r, p, "layout", "contact_new")
	}
}

func (a *App) contacts_new() http.HandlerFunc {
	// log.Printf("[contacts] contacts_new %+v\n", a.contacts.contacts)

	return func(w http.ResponseWriter, r *http.Request) {
		var c PayloadContact
		c.Errors.Record = r.ParseForm()
		if c.Errors.Record != nil {
			log.Printf("[contacts] contacts_new %v\n", c.Errors.Record)
			payload := struct {
				Contact PayloadContact
			}{
				Contact: c,
			}
			a.render(w, r, payload, "layout", "contacts/new")
			return
		}

		nc, err := a.store.CreateContact(r.FormValue("first_name"), r.FormValue("last_name"), r.FormValue("phone"), r.FormValue("email"))
		if err != nil {
			payload := struct {
				Contact PayloadContact
				Flash   string
			}{
				Contact: PayloadContact{
					Id:    nc.Id,
					First: nc.First,
					Last:  nc.Last,
					Phone: nc.Phone,
					Email: nc.Email,
				},
			}
			log.Printf("p.contact %+v\n", payload.Contact)
			a.render(w, r, payload, "layout", "contacts/new")
			return
		}

		http.SetCookie(w, &http.Cookie{
			Path:     "/",
			Name:     "flash",
			Value:    "Created New Contact!",
			HttpOnly: true,
			Secure:   true,
		})
		http.Redirect(w, r, "/contacts", http.StatusSeeOther)
	}
}

func (a *App) contacts_view() http.HandlerFunc {
	// log.Printf("[contacts] contacts %+v\n", a.contacts.contacts)

	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}

		c, err := a.store.FetchContact(id)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}

		payload := struct {
			Search  string
			Contact PayloadContact
			Flash   string
		}{
			Contact: PayloadContact{
				Id:    c.Id,
				First: c.First,
				Last:  c.Last,
				Phone: c.Phone,
				Email: c.Email,
			},
		}

		if flash, err := r.Cookie("flash"); err == nil {
			http.SetCookie(w, &http.Cookie{
				Path:     "/",
				Name:     "flash",
				MaxAge:   -1,
				HttpOnly: true,
				Secure:   true,
			})
			payload.Flash = flash.Value
		}

		a.render(w, r, payload, "layout", "contact_show")
	}
}

func (a *App) contacts_edit_get() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}

		c, err := a.store.FetchContact(id)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}

		payload := struct {
			Contact PayloadContact
		}{
			Contact: PayloadContact{
				Id:    c.Id,
				First: c.First,
				Last:  c.Last,
				Phone: c.Phone,
				Email: c.Email,
			},
		}

		a.render(w, r, payload, "layout", "contact_edit")
	}
}

func (a *App) contacts_edit_post() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}

		c, err := a.store.FetchContact(id)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}

		if err = r.ParseForm(); err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}

		c.First = r.FormValue("first_name")
		c.Last = r.FormValue("last_name")
		c.Phone = r.FormValue("phone")
		c.Email = r.FormValue("email")

		if err = a.store.UpdateContact(c); err != nil {
			payload := struct {
				Contact PayloadContact
			}{
				Contact: PayloadContact{
					Id:    c.Id,
					First: c.First,
					Last:  c.Last,
					Phone: c.Phone,
					Email: c.Email,
				},
			}
			a.render(w, r, payload, "layout", "contact_edit")
			return
		}

		http.SetCookie(w, &http.Cookie{
			Path:     "/",
			Name:     "flash",
			Value:    "Updated Contact!",
			HttpOnly: true,
			Secure:   true,
		})
		http.Redirect(w, r, fmt.Sprintf("/contacts/%d", c.Id), http.StatusSeeOther)
	}
}

func (a *App) contacts_delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}

		c, err := a.store.FetchContact(id)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}

		if err = a.store.DeleteContact(c.Id); err == nil {
			http.SetCookie(w, &http.Cookie{
				Path:     "/",
				Name:     "flash",
				Value:    "Deleted Contact!",
				HttpOnly: true,
				Secure:   true,
			})
		}

		http.Redirect(w, r, "/contacts", http.StatusSeeOther)
	}
}
