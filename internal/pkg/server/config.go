// Copyright 2021 Quanhui Li. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package server

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/marmotedu/log"
	"github.com/spf13/viper"
)

// InsecureServingInfo holds configuration of the insecure http server.
type InsecureServingInfo struct {
	Address string
}

// CertKey contains configuration items related to certificate.
type CertKey struct {
	// CertFile is a file containing a PEM-encoded certificate, and possibly the complete certificate chain
	CertFile string
	// KeyFile is a file containing a PEM-encoded private key for the certificate specified by CertFile
	KeyFile string
}

// SecureServingInfo holds configuration of the TLS server.
type SecureServingInfo struct {
	BindAddress string
	BindPort    int
	CertKey     CertKey
}

type JwtInfo struct {
	// defaults to "filestore jwt"
	Realm string
	// defaults to empty
	Key string
	// defaults to one hour
	Timeout time.Duration
	// defaults to zero
	MaxRefresh time.Duration
}

type Config struct {
	InsecureServing *InsecureServingInfo
	SecureServing   *SecureServingInfo
	Jwt             *JwtInfo
	Mode            string
	Middlewares     []string
	Healthz         bool
	EnableProfiling bool
	EnableMetrics   bool
}

func NewConfig() *Config {
	return &Config{
		Healthz:         true,
		Mode:            gin.ReleaseMode,
		Middlewares:     []string{},
		EnableProfiling: true,
		EnableMetrics:   true,
		Jwt: &JwtInfo{
			Realm:      "filestore jwt",
			Timeout:    1 * time.Hour,
			MaxRefresh: 1 * time.Hour,
		},
	}
}

// New returns a new instance of GenericAPIServer form the given config.
func (c *Config) New() (*GenericAPIServer, error) {
	s := &GenericAPIServer{
		InsecureServingInfo: c.InsecureServing,
		mode:                c.Mode,
		healthz:             c.Healthz,
		middlewares:         c.Middlewares,
		Engine:              gin.New(),
	}

	initGenericAPIServer(s)

	return s, nil
}

// LoadConfig reads in config file and ENV variables if set.
func LoadConfig(cfg string, defaultName string) {
	if cfg != "" {
		viper.SetConfigFile(cfg)
	} else {
		viper.AddConfigPath("$HOME/workspace/filestore/config")
		viper.AddConfigPath(".")
		viper.SetConfigName("filestore")
	}

	// Use config file from the flag.
	viper.SetConfigType("yaml")
	viper.AutomaticEnv()

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err != nil {
		log.Warnf("WARNING: viper failed to discover and load the configuration file: %s", err.Error())
	}
}
