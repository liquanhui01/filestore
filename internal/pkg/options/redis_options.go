// Copyright 2021 Quanhui Li. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package options

import (
	"github.com/gomodule/redigo/redis"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"

	"filestore/pkg/db"
)

// RedisOptions defines options for redis cluster.
type RedisOptions struct {
	Host                  string   `json:"host"`
	Port                  int      `json:"port"`
	Addrs                 []string `json:"addrs"`
	Username              string   `json:"username"`
	Password              string   `json:"password"`
	Database              int      `json:"database"`
	MasterName            string   `json:"mastername"`
	MaxIdle               int      `json:"maxidle"`
	MaxActive             int      `json:"maxactive"`
	Timeout               int      `json:"timeout"`
	EnableCluster         bool     `json:"enablecluster"`
	UseSSL                bool     `json:"usessl"`
	SSLInsecureSkipVerify bool     `json:"sslinsecureskipverify"`
}

// NewRedisOptions create a `zero` value instance.
func NewRedisOptions() *RedisOptions {
	return &RedisOptions{
		// Host:                  "127.0.0.1",
		// Port:                  6379,
		// Addrs:                 []string{},
		// Username:              "",
		// Password:              "",
		// Database:              0,
		// MasterName:            "",
		// MaxIdle:               2000,
		// MaxActive:             4000,
		// Timeout:               0,
		// EnableCluster:         false,
		// UseSSL:                false,
		// SSLInsecureSkipVerify: false,
		Host:                  viper.GetString("redis.host"),
		Port:                  viper.GetInt("redis.port"),
		Addrs:                 []string{},
		Username:              viper.GetString("redis.username"),
		Password:              viper.GetString("redis.password"),
		Database:              viper.GetInt("redis.database"),
		MasterName:            viper.GetString("redis.mastername"),
		MaxIdle:               viper.GetInt("redis.maxidle"),
		MaxActive:             viper.GetInt("redis.maxactive"),
		Timeout:               viper.GetInt("redis.timeout"),
		EnableCluster:         viper.GetBool("redis.enablecluster"),
		UseSSL:                viper.GetBool("redis.usessl"),
		SSLInsecureSkipVerify: viper.GetBool("redis.sslinsecureskipverify"),
	}
}

// Validate verifies flags passed to RedisOptions.
func (o *RedisOptions) Validate() []error {
	errs := []error{}

	return errs
}

// AddFlags adds flags related to redis storage for a specific APIServer to the specified FlagSet.
func (o *RedisOptions) AddFlags(fs *pflag.FlagSet) {
	fs.StringVar(&o.Host, "redis.host", o.Host, "Hostname of your Redis server.")
	fs.IntVar(&o.Port, "redis.port", o.Port, "The port the Redis server is listening on.")
	fs.StringSliceVar(&o.Addrs, "redis.addrs", o.Addrs, "A set of redis address(format: 127.0.0.1:6379).")
	fs.StringVar(&o.Username, "redis.username", o.Username, "Username for access to redis service.")
	fs.StringVar(&o.Password, "redis.password", o.Password, "Optional auth password for Redis db.")

	fs.IntVar(&o.Database, "redis.database", o.Database, ""+
		"By default, the database is 0. Setting the database is not supported with redis cluster. "+
		"As such, if you have --redis.enable-cluster=true, then this value should be omitted or explicitly set to 0.")

	fs.StringVar(&o.MasterName, "redis.mastername", o.MasterName, "The name of master redis instance.")

	fs.IntVar(&o.MaxIdle, "redis.maxidle", o.MaxIdle, ""+
		"This setting will configure how many connections are maintained in the pool when idle (no traffic). "+
		"Set the --redis.optimisation-max-active to something large, we usually leave it at around 2000 for "+
		"HA deployments.")

	fs.IntVar(&o.MaxActive, "redis.maxactive", o.MaxActive, ""+
		"In order to not over commit connections to the Redis server, we may limit the total "+
		"number of active connections to Redis. We recommend for production use to set this to around 4000.")

	fs.IntVar(&o.Timeout, "redis.timeout", o.Timeout, "Timeout (in seconds) when connecting to redis service.")

	fs.BoolVar(&o.EnableCluster, "redis.enablecluster", o.EnableCluster, ""+
		"If you are using Redis cluster, enable it here to enable the slots mode.")

	fs.BoolVar(&o.UseSSL, "redis.usessl", o.UseSSL, ""+
		"If set, IAM will assume the connection to Redis is encrypted. "+
		"(use with Redis providers that support in-transit encryption).")

	fs.BoolVar(&o.SSLInsecureSkipVerify, "redis.sslinsecureskipverify", o.SSLInsecureSkipVerify, ""+
		"Allows usage of self-signed certificates when connecting to an encrypted Redis database.")
}

func (r *RedisOptions) NewClient() *redis.Pool {
	opts := &db.RedisOptions{
		Host:                  r.Host,
		Port:                  r.Port,
		Addrs:                 r.Addrs,
		Username:              r.Username,
		Password:              r.Password,
		Database:              r.Database,
		MasterName:            r.MasterName,
		MaxIdle:               r.MaxIdle,
		MaxActive:             r.MaxActive,
		Timeout:               r.Timeout,
		EnableCluster:         r.EnableCluster,
		UseSSL:                r.UseSSL,
		SSLInsecureSkipVerify: r.SSLInsecureSkipVerify,
	}
	return db.NewRedis(opts)
}
