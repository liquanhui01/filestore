// Copyright 2021 Quanhui Li. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package config

import (
	"log"
	"sync"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

var (
	cfg  = pflag.StringP("conf", "c", "", "configuration file.")
	help = pflag.BoolP("help", "h", false, "Show this help message.")
)

func init() {
	var once sync.Once
	once.Do(func() {
		pflag.Parse()
		if *help {
			pflag.Usage()
		}

		// Read configuration from file.
		if *cfg != "" {
			viper.SetConfigFile(*cfg)
			viper.SetConfigType("yaml")
		} else {
			viper.AddConfigPath(".")
			viper.AddConfigPath("$HOME/workspace/filestore/config")
			viper.SetConfigName("filestore")
			viper.SetConfigType("yaml")
		}

		if err := viper.ReadInConfig(); err != nil {
			log.Fatal("failed to initlize configuration.")
		}
	})
}
