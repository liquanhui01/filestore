// Copyright 2021 Quanhui Li. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package config

import (
	"log"
	"sync"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"

	"github.com/liquanhui01/filestore/internal/apiserver/options"
)

var (
	cfg  = pflag.StringP("conf", "c", "", "configuration file.")
	help = pflag.BoolP("help", "h", false, "Show this help message.")
)

type Config struct {
	*options.Options
}

// CreateConfigFromOptions creates a running configuration instance based
// on a given filestore command line or configuration file option.
func CreateConfigFromOptions(opts *options.Options) (*Config, error) {
	return &Config{opts}, nil
}

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
		} else {
			viper.AddConfigPath(".")
			viper.AddConfigPath("$HOME/workspace/filestore/config")
			viper.AddConfigPath("$HOME/filestore/config")
			viper.SetConfigName("filestore")
		}
		viper.SetConfigType("yaml")

		// Read in config
		if err := viper.ReadInConfig(); err != nil {
			log.Fatalf("Read config failed, err is: %s\n", err)
		}
	})
}
