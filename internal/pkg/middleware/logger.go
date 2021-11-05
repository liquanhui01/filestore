// Copyright 2021 Quanhui Li. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package middleware

import (
	"fmt"
	"os"
	"path"
	"time"

	"github.com/gin-gonic/gin"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
)

var logger *logrus.Logger

type LoggerOptions struct {
	path         string
	filename     string
	level        logrus.Level
	maxAge       int
	rotationTime int
}

// NewLogger init
func NewLogger(opts *LoggerOptions) *LoggerOptions {
	return &LoggerOptions{
		path:         opts.path,
		filename:     opts.filename,
		level:        opts.level,
		maxAge:       opts.maxAge,
		rotationTime: opts.rotationTime,
	}
}

// initConfig configs logrus.
func (l *LoggerOptions) Config() *logrus.Logger {
	fileName := path.Join(l.path, l.filename)

	// Write in file
	f, err := os.OpenFile(fileName, os.O_WRONLY|os.O_APPEND, os.ModeAppend)
	if err != nil {
		fmt.Printf("Logger middleware error, err is: %s\n", err.Error())
	}

	// init
	logger = logrus.New()
	// set out pattern
	logger.Out = f
	// set out level
	logger.SetLevel(l.level)
	// set rotatelog
	logWriter, _ := rotatelogs.New(
		fileName+"%Y%m%d.log",
		rotatelogs.WithLinkName(fileName),
		rotatelogs.WithMaxAge(time.Duration(l.maxAge)*time.Hour),
		rotatelogs.WithRotationTime(time.Duration(l.rotationTime)*time.Hour),
	)

	writeMap := lfshook.WriterMap{
		logrus.InfoLevel:  logWriter,
		logrus.FatalLevel: logWriter,
		logrus.DebugLevel: logWriter,
		logrus.WarnLevel:  logWriter,
		logrus.ErrorLevel: logWriter,
		logrus.PanicLevel: logWriter,
	}

	lfhook := lfshook.NewHook(writeMap, &logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})
	// 新增hook
	logger.AddHook(lfhook)

	return logger
}

// Logger middleware
func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()
		// 处理请求
		c.Next()
		endTime := time.Now()
		// time
		latencyTime := endTime.Sub(startTime)
		// reqeust method
		reqMethod := c.Request.Method
		// request router
		reqUrl := c.Request.URL
		// status code
		statuCode := c.Writer.Status()
		// request client ip
		clientIP := c.ClientIP()

		logger.Infof("| %3d | %13v | %15s | %s | %s",
			statuCode,
			latencyTime,
			clientIP,
			reqMethod,
			reqUrl,
		)
	}
}
