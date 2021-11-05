// Copyright 2021 Quanhui Li. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

// apiserver is the api server for filestore service.
package main

import (
	"math/rand"
	"os"
	"runtime"
	"time"

	"github.com/liquanhui01/filestore/internal/apiserver"
	_ "github.com/liquanhui01/filestore/internal/apiserver/config"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	if len(os.Getenv("GOMAXPROCS")) == 0 {
		runtime.GOMAXPROCS(runtime.NumCPU())
	}

	// start the server
	apiserver.NewApp()
}
