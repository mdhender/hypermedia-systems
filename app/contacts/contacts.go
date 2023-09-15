// Copyright (c) 2023 Michael D Henderson. All rights reserved.

package contacts

import (
	"fmt"
	"io"
	"strings"
)

type Contacts struct {
	contacts []*Contact
}

func (c *Contacts) All() *Contacts {
	if c == nil {
		return &Contacts{}
	}
	return c
}

func (c *Contacts) Dump(w io.Writer) {
	if c == nil {
		_, _ = fmt.Fprintf(w, "[contacts] dump: nil\n")
		return
	} else if len(c.contacts) == 0 {
		_, _ = fmt.Fprintf(w, "[contacts] dump: []\n")
		return
	}
	for i, contact := range c.contacts {
		_, _ = fmt.Fprintf(w, "[contacts] dump: %d %+v\n", i+1, *contact)
	}
}

func (c *Contacts) Search(search string) *Contacts {
	s := &Contacts{}
	if c == nil {
		return s
	}
	matches := func(a, b string) bool {
		return strings.EqualFold(a, b)
	}
	for _, contact := range c.contacts {
		if matches(search, contact.First) {
			s.contacts = append(s.contacts, contact)
		} else if matches(search, contact.Last) {
			s.contacts = append(s.contacts, contact)
		} else if matches(search, contact.Phone) {
			s.contacts = append(s.contacts, contact)
		} else if matches(search, contact.Email) {
			s.contacts = append(s.contacts, contact)
		}
	}
	return s
}

type Contact struct {
	Id    int
	First string
	Last  string
	Phone string
	Email string
}

func NewContacts(contacts ...*Contact) *Contacts {
	c := &Contacts{}
	for _, contact := range contacts {
		c.contacts = append(c.contacts, contact)
	}
	return c
}

func NewContact(id int, first, last, phone, email string) *Contact {
	return &Contact{
		Id:    id,
		First: first,
		Last:  last,
		Phone: phone,
		Email: email,
	}
}
