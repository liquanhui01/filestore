// Copyright 2021 Quanhui Li. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package config

import "filestore/internal/apiserver/options"

type Config struct {
	*options.Options
}

// CreateConfigFromOptions creates a running configuration instance based
// on a given filestore command line or configuration file option.
func CreateConfigFromOptions(opts *options.Options) (*Config, error) {
	return &Config{opts}, nil
}
