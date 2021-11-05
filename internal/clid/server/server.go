// Copyright 2021 Quanhui Li. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package server

import (
	"github.com/spf13/cobra"

	"github.com/liquanhui01/filestore/internal/apiserver"
)

func NewServerCmd() *cobra.Command {
	var cmd = &cobra.Command{
		Use:     "server",
		Short:   "filestore platform server.",
		Long:    "ignore",
		Aliases: []string{"srv"},
		Run: func(cmd *cobra.Command, args []string) {
			// r := gin.Default()
			// r.GET("/ping", func(c *gin.Context) {
			// 	c.JSON(200, gin.H{
			// 		"message": "pong",
			// 	})
			// })
			// r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
			apiserver.NewApp()
		},
	}

	return cmd
}
