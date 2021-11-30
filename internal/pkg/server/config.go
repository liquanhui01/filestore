// Copyright 2021 Quanhui Li. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package server

import (
	"time"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
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
		Healthz:     true,
		Mode:        viper.GetString("gin.mode"),
		Middlewares: []string{"recovery", "secret", "requestid", "logger"},
		InsecureServing: &InsecureServingInfo{
			Address: viper.GetString("server.host"),
		},
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
func (c *Config) New() {
	s := &GenericAPIServer{
		InsecureServingInfo: c.InsecureServing,
		Mode:                c.Mode,
		Healthz:             c.Healthz,
		Middlewares:         c.Middlewares,
		Engine:              gin.New(),
	}

	InitGenericAPIServer(s)
}

// LoadConfig reads in config file and ENV variables if set.
func LoadConfig(cfg string, defaultName string) {
	if cfg != "" {
		viper.SetConfigFile(cfg)
	} else {
		viper.AddConfigPath("$HOME/filestore/config")
		viper.AddConfigPath(".")
		viper.AddConfigPath("$HOME/workspace/filestore/config")
		viper.SetConfigName(defaultName)
	}

	viper.SetConfigType("yaml")
	viper.AutomaticEnv()

	// Load config file
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Failed to discover and load configuration file, err is %s\n", err.Error())
	}
}
