// Copyright (c) 2023 Michael D Henderson. All rights reserved.

package contacts

import (
	"fmt"
	"os"
)

type Option func(*App) error

func WithContacts(c *Contacts) Option {
	return func(a *App) error {
		if a.contacts == nil {
			a.contacts = c
		} else {
			for _, contact := range c.contacts {
				a.contacts.contacts = append(a.contacts.contacts, contact)
			}
		}
		return nil
	}
}

func WithTemplates(path string) Option {
	return func(a *App) error {
		if sb, err := os.Stat(path); err != nil {
			return err
		} else if !sb.IsDir() {
			return fmt.Errorf("invalid path %q", path)
		}
		a.templates = path
		return nil
	}
}
