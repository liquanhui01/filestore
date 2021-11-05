// Copyright 2021 Quanhui Li. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package clid

import (
	"github.com/spf13/cobra"

	srv "github.com/liquanhui01/filestore/internal/clid/server"
	ver "github.com/liquanhui01/filestore/internal/clid/version"
)

func NewDefaultCommand() *cobra.Command {
	var cmds = &cobra.Command{
		Use:     "filectl",
		Short:   "filectl control the filestore platform",
		Long:    `filectl controls the filestore platform, is the client side tool for filestore platform.`,
		Aliases: []string{"fctl", "filectl"},
		Run:     runHelp,
	}
	cmds.AddCommand(ver.NewVersionCmd())
	cmds.AddCommand(srv.NewServerCmd())

	return cmds
}

func runHelp(cmd *cobra.Command, args []string) {
	_ = cmd.Help()
}
