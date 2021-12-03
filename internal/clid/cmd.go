// Copyright 2021 Quanhui Li. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package clid

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	genericfileflag "github.com/liquanhui01/filestore/cmd"
	"github.com/liquanhui01/filestore/internal/apiserver/store/mysql"
	srv "github.com/liquanhui01/filestore/internal/clid/server"
	ver "github.com/liquanhui01/filestore/internal/clid/version"
	"github.com/liquanhui01/filestore/internal/pkg/options"
	"github.com/liquanhui01/filestore/internal/pkg/server"
	genericapiserver "github.com/liquanhui01/filestore/internal/pkg/server"
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

	cobra.OnInitialize(func() {
		genericapiserver.LoadConfig(viper.GetString(genericfileflag.FlagConfig), "filestore")
		_, err := mysql.GetMySQLFactoryOr(options.NewMySQLOptions())
		if err != nil {
			log.Fatal("failed to initialize database")
		}
		server.NewConfig().New()
	})

	return cmds
}

func runHelp(cmd *cobra.Command, args []string) {
	_ = cmd.Help()
}
