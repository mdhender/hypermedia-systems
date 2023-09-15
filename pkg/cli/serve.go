// Copyright (c) 2023 Michael D Henderson. All rights reserved.

package cli

import (
	"github.com/spf13/cobra"
)

var (
	argsServe = struct{}{}

	cmdServe = &cobra.Command{
		Use:   "serve",
		Short: "run application server",
		Long:  `Run an application server.`,
	}
)
