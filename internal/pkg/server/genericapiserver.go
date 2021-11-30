// Copyright 2021 Quanhui Li. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package server

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"golang.org/x/sync/errgroup"

	"github.com/liquanhui01/filestore/internal/pkg/middleware"
)

// GenericAPIServer contains state for a filestore api server.
// type GenericAPIServer gin.Engin.
type GenericAPIServer struct {
	Middlewares []string
	Mode        string
	// InsecureServingInfo holds configuration of the insecure HTTP server.
	InsecureServingInfo *InsecureServingInfo

	// SecureServingInfo holds configutation of the TLS server.
	SecureServingInfo *SecureServingInfo

	ShutdownTimeout time.Duration

	*gin.Engine
	Healthz bool

	insecureServer, secureServer *http.Server
}

func InitGenericAPIServer(c *GenericAPIServer) {
	c.Setup()
	c.InstallMiddlewares()
	c.InstallAPIs()

}

func (s *GenericAPIServer) Setup() {
	gin.SetMode(s.Mode)
}

// InstallMiddlewares install all middlewares.
func (s *GenericAPIServer) InstallMiddlewares() {
	// necessary middlewares
	s.Use(middleware.RequestID())
	s.Use(middleware.Context())
	s.Use(middleware.Logger())

	// install custom middlewares
	for _, m := range s.Middlewares {
		mv, ok := middleware.Middlewares[m]
		if !ok {
			continue
		}
		s.Use(mv)
	}
}

// InstallAPIs install generic apis.
func (s *GenericAPIServer) InstallAPIs() {
	if s.Healthz {
		s.GET("/healthz", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"msg":  "health ok",
				"data": nil,
			})
		})
	}
}

// Run spawns the http server. It only returns when the port cannot be listened on initially.
func (s *GenericAPIServer) Run() error {
	// For scalability, use custom HTTP configuration mode here
	s.insecureServer = &http.Server{
		Addr:    s.insecureServer.Addr,
		Handler: s,
	}

	// For scalability, use custom HTTP configuration mode here
	s.secureServer = &http.Server{
		Addr:    s.secureServer.Addr,
		Handler: s,
	}

	var eg errgroup.Group

	// Initializing the server in a goroutine so that
	// it won't block thr graceful shutdown handling below.
	eg.Go(func() error {
		log.Infof("Start to listening the incoming requests on http address: %s\n", s.InsecureServingInfo.Address)

		if err := s.insecureServer.ListenAndServe(); err != nil && errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("Failed to listening the insecure server, err is %s\n", err.Error())
			return err
		}

		log.Infof("Server on %s stopped", s.InsecureServingInfo.Address)

		return nil
	})

	eg.Go(func() error {
		key, cert := s.SecureServingInfo.CertKey.KeyFile, s.SecureServingInfo.CertKey.CertFile
		if key == "" || cert == "" || s.SecureServingInfo.BindPort == 0 {
			return nil
		}

		log.Infof("Start to listening the incoming requests on https address: %s", s.SecureServingInfo.BindAddress)

		if err := s.secureServer.ListenAndServeTLS(cert, key); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("Failed to listening the secure server, err is %s\n", err.Error())

			return err
		}

		log.Infof("Server on %s stopped", s.SecureServingInfo.BindAddress)

		return nil
	})

	// Ping the server to make sure thr router is running.
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if s.Healthz {
		if err := s.ping(ctx); err != nil {
			return err
		}
	}

	if err := eg.Wait(); err != nil {
		log.Fatalf(err.Error())
	}

	return nil
}

// Close graceful shutdown the api server.
func (s *GenericAPIServer) Close() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := s.insecureServer.Shutdown(ctx); err != nil {
		log.Warnf("Shutdown insecure server failed: %s", err.Error())
	}

	if err := s.secureServer.Shutdown(ctx); err != nil {
		log.Warnf("Shutdown secure server failed: %s", err.Error())
	}
}

func (s *GenericAPIServer) ping(ctx context.Context) error {
	url := fmt.Sprintf("http://%s/healthz", s.InsecureServingInfo.Address)
	if strings.Contains(s.InsecureServingInfo.Address, "0.0.0.0") {
		url = fmt.Sprintf("http://127.0.0.1:%s/healthz", strings.Split(s.InsecureServingInfo.Address, ":")[1])
	}

	for {
		// Change NewRequest to NewRequestWithContext and pass context it.
		req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
		if err != nil {
			return err
		}

		resp, err := http.DefaultClient.Do(req)
		if err == nil && resp.StatusCode == http.StatusOK {
			log.Info("The router has been deployed successfully.")

			resp.Body.Close()

			return nil
		}

		// Sleep for a second to continue the next ping.
		log.Infof("Waiting for the router, retry in 1 second.")
		time.Sleep(1 * time.Second)

		select {
		case <-ctx.Done():
			log.Fatalf("can not ping http server within the specified time interval.")
		default:
		}
	}
}
