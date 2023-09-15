// Copyright (c) 2023 Michael D Henderson. All rights reserved.

package cli

import (
	"github.com/go-chi/chi/middleware"
	"github.com/mdhender/hypermedia-systems/app/contacts"
	"github.com/mdhender/hypermedia-systems/pkg/mw"
	"github.com/spf13/cobra"
	"log"
	"path/filepath"
)

var (
	argsServeContacts = struct {
		Host       string
		Port       string
		Middleware struct {
			BadRunes bool
			CORS     bool
			Logging  bool
		}
		Templates string // path to templates
	}{
		Port:      "8080",
		Templates: ".",
	}

	cmdServeContacts = &cobra.Command{
		Use:   "contacts",
		Short: "run contacts app",
		Long:  `Run the contacts application.`,
		Run: func(cmd *cobra.Command, args []string) {
			path, err := filepath.Abs(argsServeContacts.Templates)
			if err != nil {
				log.Fatalf("[contacts] templates: %v\n", err)
			}
			argsServeContacts.Templates = path

			log.Printf("[contacts] serve contacts %+v\n", argsServeContacts)

			var options []contacts.Option
			options = append(options, contacts.WithTemplates(argsServeContacts.Templates))
			options = append(options, contacts.WithHost(argsServeContacts.Host))
			options = append(options, contacts.WithPort(argsServeContacts.Port))
			if argsServeContacts.Middleware.BadRunes {
				options = append(options, contacts.WithBadRunesMiddleware(mw.BadRunes))
			}
			if argsServeContacts.Middleware.CORS {
				options = append(options, contacts.WithCorsMiddleware(mw.CORS))
			}
			if argsServeContacts.Middleware.Logging {
				options = append(options, contacts.WithLoggingMiddleware(middleware.Logger))
			}

			options = append(options, contacts.WithContacts(contacts.NewContacts(
				contacts.NewContact(42, "John", "Smith", "303/555.2345", "john@example.com"),
				contacts.NewContact(43, "Dana", "Crandith", "303/555.1212", "dcran@example.com"),
				contacts.NewContact(44, "Edith", "Neutvaar", "303/555.9876", "en@example.com"),
			)))

			app, err := contacts.New(options...)
			if err != nil {
				log.Fatalf("[contacts] app: %v\n", err)
			}

			if err := app.ListenAndServe(); err != nil {
				log.Fatal(err)
			}
		},
	}
)
