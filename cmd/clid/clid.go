// Copyright 2021 Quanhui Li. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package main

import (
	"os"

	"github.com/liquanhui01/filestore/internal/clid"
)

func main() {
	command := clid.NewDefaultCommand()
	if err := command.Execute(); err != nil {
		os.Exit(1)
	}
}
