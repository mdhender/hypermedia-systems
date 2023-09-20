// Copyright (c) 2023 Michael D Henderson. All rights reserved.

package store

import "strings"

type Contact struct {
	Id    int
	First string
	Last  string
	Phone string
	Email string
}

func (s *Store) nextID() int {
	s.Lock()
	defer s.Unlock()

	id := s.nextId
	s.nextId++

	return id
}

func (s *Store) CreateContact(first, last, phone, email string) (Contact, error) {
	contact := Contact{
		Id:    s.nextID(),
		First: first,
		Last:  last,
		Phone: phone,
		Email: email,
	}

	s.Lock()
	defer s.Unlock()

	s.contacts = append(s.contacts, &contact)

	return contact, nil
}

func (s *Store) DeleteContact(id int) error {
	s.Lock()
	defer s.Unlock()

	var t []*Contact
	found := false
	for _, c := range s.contacts {
		if c.Id == id {
			found = true
			continue
		}
		t = append(t, c)
	}
	s.contacts = t

	if !found {
		return ErrNotFound
	}
	return nil
}

func (s *Store) FetchContact(id int) (Contact, error) {
	s.Lock()
	defer s.Unlock()

	for _, c := range s.contacts {
		if c.Id == id {
			return *c, nil
		}
	}

	return Contact{}, ErrNotFound
}

func (s *Store) UpdateContact(nc Contact) error {
	s.Lock()
	defer s.Unlock()

	for _, c := range s.contacts {
		if nc.Id == c.Id {
			c.First = nc.First
			c.Last = nc.Last
			c.Phone = nc.Phone
			c.Email = nc.Email
			return nil
		}
	}

	return ErrNotFound
}

func (s *Store) AllContacts() []Contact {
	s.Lock()
	defer s.Unlock()

	var cc []Contact
	for _, c := range s.contacts {
		cc = append(cc, *c)
	}
	return cc
}

func (s *Store) SearchContacts(search string) []Contact {
	s.Lock()
	defer s.Unlock()

	matches := func(a, b string) bool {
		return strings.EqualFold(a, b)
	}

	var cc []Contact
	for _, c := range s.contacts {
		if matches(search, c.First) {
			cc = append(cc, *c)
		} else if matches(search, c.Last) {
			cc = append(cc, *c)
		} else if matches(search, c.Phone) {
			cc = append(cc, *c)
		} else if matches(search, c.Email) {
			cc = append(cc, *c)
		}
	}

	return cc
}
