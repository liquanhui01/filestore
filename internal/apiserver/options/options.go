// Copyright 2021 Quanhui Li. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

// Package options contains flags and options for initializing an apiserver
package options

import (
	genericoptions "github.com/liquanhui01/filestore/internal/pkg/options"
)

type Options struct {
	MySQLOptions *genericoptions.MySQLOptions
	RedisOptions *genericoptions.RedisOptions
}

// NewOptions creates a new Options object with default parameters.
func NewOptions() *Options {
	o := &Options{
		MySQLOptions: genericoptions.NewMySQLOptions(),
		RedisOptions: genericoptions.NewRedisOptions(),
	}
	return o
}
