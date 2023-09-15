// Copyright (c) 2023 Michael D Henderson. All rights reserved.

package cli

import (
	"github.com/mdhender/hypermedia-systems/app/contacts"
	"github.com/mdhender/hypermedia-systems/server"
	"github.com/spf13/cobra"
	"log"
	"path/filepath"
)

var (
	argsServeContacts = struct {
		Templates string // path to templates
	}{
		Templates: ".",
	}

	cmdServeContacts = &cobra.Command{
		Use:   "contacts",
		Short: "run contacts server",
		Long:  `Run the contacts application.`,
		Run: func(cmd *cobra.Command, args []string) {
			path, err := filepath.Abs(argsServeContacts.Templates)
			if err != nil {
				log.Fatalf("[contacts] templates: %v\n", err)
			}
			argsServeContacts.Templates = path

			log.Printf("[contacts] serve contacts %+v\n", argsServeContacts)

			var contactOptions []contacts.Option
			contactOptions = append(contactOptions, contacts.WithTemplates(argsServeContacts.Templates))
			contactOptions = append(contactOptions, contacts.WithContacts(contacts.NewContacts(
				contacts.NewContact(42, "John", "Smith", "303/555.2345", "john@example.com"),
				contacts.NewContact(43, "Dana", "Crandith", "303/555.1212", "dcran@example.com"),
				contacts.NewContact(44, "Edith", "Neutvaar", "303/555.9876", "en@example.com"),
			)))

			app, err := contacts.New(contactOptions...)
			if err != nil {
				log.Fatalf("[contacts] app: %v\n", err)
			}

			log.Printf("[contacts] serve %+v\n", argsServe)

			serverOptions := []server.Option{server.WithApplication(app)}
			if argsServe.Host != "" {
				serverOptions = append(serverOptions, server.WithHost(argsServe.Host))
			}
			serverOptions = append(serverOptions, server.WithPort(argsServe.Port))
			if argsServe.BadRunesMiddleware {
				serverOptions = append(serverOptions, server.WithBadRunesMiddleware())
			}
			if argsServe.CORSMiddleware {
				serverOptions = append(serverOptions, server.WithCorsMiddleware())
			}
			s, err := server.New(serverOptions...)
			if err != nil {
				log.Fatalf("[contacts] %v", err)
			} else if s.Handler == nil {
				log.Fatalf("[contacts] missing handler\n")
			}

			if err := s.Serve(); err != nil {
				log.Fatalf("[contacts] %v", err)
			}
		},
	}
)
