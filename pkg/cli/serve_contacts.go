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
		Long:  `Run an contacts server.`,
		Run: func(cmd *cobra.Command, args []string) {
			path, err := filepath.Abs(argsServeContacts.Templates)
			if err != nil {
				log.Fatalf("[contacts] templates: %v\n", err)
			}
			argsServeContacts.Templates = path

			log.Printf("[contacts] serve contacts %+v\n", argsServeContacts)

			var contactsOptions []contacts.Option
			contactsOptions = append(contactsOptions, contacts.WithTemplates(argsServeContacts.Templates))
			contactsOptions = append(contactsOptions, contacts.WithContacts(contacts.NewContacts(
				contacts.NewContact("jim", "Jim"),
				contacts.NewContact("joe", "Joe"),
			)))

			app, err := contacts.New(contactsOptions...)
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
				log.Fatalf("[serve] %v", err)
			} else if s.Handler == nil {
				log.Fatalf("[serve] missing handler\n")
			}

			if err := s.Serve(); err != nil {
				log.Fatalf("[serve] %v", err)
			}
		},
	}
)
