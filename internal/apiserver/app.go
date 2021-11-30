// Copyright 2021 Quanhui Li. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package apiserver

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"golang.org/x/sync/errgroup"
)

// NewApp defines a server app
func NewApp() {

	gin.SetMode(viper.GetString("gin.mode"))

	addr := viper.GetString("server.host")
	router := InitRouter()
	srv := &http.Server{
		Addr:    addr,
		Handler: router,
	}

	var eg errgroup.Group
	eg.Go(func() error {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("failed to run server")
			return err
		}

		return nil
	})

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server shutdown:%s\n", err)
	}

	if err := eg.Wait(); err != nil {
		log.Fatalf(err.Error())
	}

	log.Fatalf("Server exiting")
}
