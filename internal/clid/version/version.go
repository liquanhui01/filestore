// Copyright 2021 Quanhui Li. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package version

import (
	"fmt"

	"github.com/spf13/cobra"
)

func NewVersionCmd() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "version",
		Short: "Print the client and server version information",
		Long:  "Print the client and server version infomation for the current context",
		Example: `# Print the client and server versions for the current context
		filestore version`,
		Run: func(cmd *cobra.Command, args []string) {
			// TODO
			fmt.Printf("version is %s\n", "v1.0.0")
		},
	}

	return cmd
}
