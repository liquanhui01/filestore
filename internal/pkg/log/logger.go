// Copyright 2021 Quanhui Li. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package log

import "github.com/sirupsen/logrus"

type LoggerOptions struct {
	path     string
	filename string
	level    logrus.Level
}

// NewLogger init
func NewLogger(opts *LoggerOptions) *LoggerOptions {
	return &LoggerOptions{
		path:     opts.path,
		filename: opts.filename,
		level:    opts.level,
	}
}
