// Copyright 2021 Quanhui Li. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package options

import (
	"time"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"gorm.io/gorm"

	"github.com/liquanhui01/filestore/pkg/db"
)

// MySQLOptions defines options for mysql database.
type MySQLOptions struct {
	Host                  string        `json:"host,omitempty"`
	Username              string        `json:"username,omitempty"`
	Password              string        `json:"-"`
	Database              string        `json:"database"`
	MaxIdleConnections    int           `json:"maxidleconnections,omitempty"`
	MaxOpenConnections    int           `json:"maxopenconnections,omitempty"`
	MaxConnectionLifeTime time.Duration `json:"maxconnectionlifetime,omitempty"`
}

// NewMySQLOptions create a `zero` value instance.
func NewMySQLOptions() *MySQLOptions {
	return &MySQLOptions{
		// Host:                  "127.0.0.1:3306",
		// Username:              "",
		// Password:              "",
		// Database:              "",
		// MaxIdleConnections:    100,
		// MaxOpenConnections:    100,
		// MaxConnectionLifeTime: time.Duration(10) * time.Second,
		Host:                  viper.GetString("mysql.host"),
		Username:              viper.GetString("mysql.username"),
		Password:              viper.GetString("mysql.password"),
		Database:              viper.GetString("mysql.database"),
		MaxIdleConnections:    viper.GetInt("mysql.maxidleconnections"),
		MaxOpenConnections:    viper.GetInt("mysql.maxopenconnections"),
		MaxConnectionLifeTime: time.Duration(viper.GetDuration("mysql.maxconnectionlifetime")) * time.Second,
	}
}

// Validate verifies flags passed to MySQLOptions.
func (o *MySQLOptions) Validate() []error {
	errs := []error{}

	return errs
}

// AddFlags adds flags related to mysql for a specific APIServer to the specified FlagSet.
func (o *MySQLOptions) AddFlags(fs *pflag.FlagSet) {
	fs.StringVar(&o.Host, "mysql.host", o.Host, ""+
		"MySQL serviec host address. If left blank, the following related mysql options will be ignored.")

	fs.StringVar(&o.Username, "mysql.username", o.Username, ""+
		"Username for access to mysql service.")

	fs.StringVar(&o.Password, "mysql.password", o.Password, ""+
		"Password for access to mysql, should be used pair with password.")

	fs.StringVar(&o.Database, "mysql.database", o.Database, ""+
		"Database name for the server to use.")

	fs.IntVar(&o.MaxIdleConnections, "mysql.maxidleconnections", o.MaxOpenConnections, ""+
		"Maximum idle connections allowed to connect to mysql.")

	fs.IntVar(&o.MaxOpenConnections, "mysql.maxopenconnections", o.MaxOpenConnections, ""+
		"Maximum open connections allowed to connect to mysql.")

	fs.DurationVar(&o.MaxConnectionLifeTime, "mysql.maxconnectionlifetime", o.MaxConnectionLifeTime, ""+
		"Maximum connection life time allowed to connecto to mysql.")
}

// NewClient create mysql store with the given config.
func (o *MySQLOptions) NewClient() (*gorm.DB, error) {
	opts := &db.MysQLOptions{
		Host:                  o.Host,
		Username:              o.Username,
		Password:              o.Password,
		Database:              o.Database,
		MaxIdleConnections:    o.MaxIdleConnections,
		MaxOpenConnections:    o.MaxOpenConnections,
		MaxConnectionLifeTime: o.MaxConnectionLifeTime,
	}
	return db.New(opts)
}
