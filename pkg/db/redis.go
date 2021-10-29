// Copyright 2021 Quanhui Li. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package db

import (
	"fmt"
	"strconv"
	"time"

	"github.com/gomodule/redigo/redis"
)

var pool *redis.Pool

// RedisOptions defines options for redis cluster.
type RedisOptions struct {
	Host                  string
	Port                  int
	Addrs                 []string
	Username              string
	Password              string
	Database              int
	MasterName            string
	MaxIdle               int
	MaxActive             int
	Timeout               int
	EnableCluster         bool
	UseSSL                bool
	SSLInsecureSkipVerify bool
}

// NewRedisPool create the poll for redis connection.
func newRedisPool(opts *RedisOptions) *redis.Pool {
	fmt.Printf("maxActive:%d\n", opts.Port)
	address := opts.Host + ":" + strconv.Itoa(opts.Port)
	return &redis.Pool{
		MaxIdle:     opts.MaxIdle,
		MaxActive:   opts.MaxActive,
		IdleTimeout: time.Duration(opts.Timeout) * time.Second,
		Dial: func() (redis.Conn, error) {
			// Open connection
			conn, err := redis.Dial("tcp", address)
			if err != nil {
				fmt.Println("错误为:", err)
				return nil, err
			}

			// Access Certification
			if _, err := conn.Do("AUTH", "alice", "p1pp0"); err != nil {
				fmt.Println("Do err: ", err)
				conn.Close()
				return nil, err
			}

			return conn, nil
		},
		TestOnBorrow: func(conn redis.Conn, t time.Time) error {
			if time.Since(t) < time.Minute {
				return nil
			}

			_, err := conn.Do("PING")
			if err != nil {
				return err
			}

			return nil
		},
	}
}

func NewRedis(opts *RedisOptions) *redis.Pool {
	pool = newRedisPool(opts)
	return pool
}
