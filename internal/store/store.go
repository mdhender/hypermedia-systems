// Copyright (c) 2023 Michael D Henderson. All rights reserved.

package store

import (
	"fmt"
	"sync"
)

type Store struct {
	sync.Mutex

	nextId   int
	contacts []*Contact
}

var (
	ErrNotFound = fmt.Errorf("not found")
)

func New() *Store {
	s := &Store{
		nextId: 42,
	}

	s.CreateContact("John", "Smith", "303/555.2345", "john@example.com")
	s.CreateContact("Dana", "Crandith", "303/555.1212", "dcran@example.com")
	s.CreateContact("Edith", "Neutvaar", "303/555.9876", "en@example.com")

	return s
}
