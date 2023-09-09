// Copyright (c) 2023 Michael D Henderson. All rights reserved.

package flask

type Contacts struct {
	contacts []*Contact
}

func (c *Contacts) All() *Contacts {
	if c == nil {
		return &Contacts{}
	}
	return c
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
