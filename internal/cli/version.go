// Copyright (c) 2023 Michael D Henderson. All rights reserved.

package cli

import (
	"fmt"
	"github.com/mdhender/hypermedia-systems/internal/semver"
	"github.com/spf13/cobra"
)

var (
	argsVersion = semver.Version{Major: 0, Minor: 1}

	cmdVersion = &cobra.Command{
		Use:   "version",
		Short: "show application version",
		Long:  `Show application version.`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("%s\n", argsVersion)
		},
	}
)
