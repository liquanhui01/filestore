// Copyright 2021 Quanhui Li. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package server

import (
	"os"
	"os/signal"
)

var onlyOneSignalHandler = make(chan struct{})

var shutdoenHandler chan os.Signal

// SetupSignalHandler registered for SIGTERM and SIGINT. A stop channel is returned
// which is closed on one of these signals. If a second signal is caught, the program
// is terminated with exit code 1.
func SetupSignalHandler() <-chan struct{} {
	close(onlyOneSignalHandler)

	shutdoenHandler = make(chan os.Signal, 2)

	stop := make(chan struct{})

	signal.Notify(shutdoenHandler, shutdownSignals...)

	go func() {
		<-shutdoenHandler
		close(stop)
		<-shutdoenHandler
		os.Exit(1) // second signal. Exit directly.
	}()

	return stop
}

// RequestShutdown emulates a received event that is considered as shutdown signal (SIGTERM/SIGINT)
// This returns whether a handler was notified.
func RequestShutdown() bool {
	if shutdoenHandler != nil {
		select {
		case shutdoenHandler <- shutdownSignals[0]:
			return true
		default:
		}
	}

	return false
}
