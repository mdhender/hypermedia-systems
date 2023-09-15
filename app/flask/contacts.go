// Copyright (c) 2023 Michael D Henderson. All rights reserved.

package flask

import (
	"fmt"
	"io"
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

func (c *Contacts) Search(search ...string) *Contacts {
	s := &Contacts{}
	if c == nil {
		return s
	}
	for _, id := range search {
		for _, contact := range c.contacts {
			if id == contact.Id {
				s.contacts = append(s.contacts, contact)
			}
		}
	}
	return s
}

type Contact struct {
	Id   string
	Name string
}

func NewContacts(contacts ...*Contact) *Contacts {
	c := &Contacts{}
	for _, contact := range contacts {
		c.contacts = append(c.contacts, contact)
	}
	return c
}

func NewContact(id, name string) *Contact {
	return &Contact{Id: id, Name: name}
}
