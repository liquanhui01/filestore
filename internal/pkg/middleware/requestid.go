// Copyright 2021 Quanhui Li. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package middleware

import (
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

const (
	XRequestIDKey = "X-Request-ID"
)

// RequestID is a middleware that injects a 'X-Request-ID' into the context and request/response header of each request.
func RequestID() gin.HandlerFunc {
	return func(c *gin.Context) {
		rid := c.GetHeader(XRequestIDKey)

		if rid == "" {
			rid = uuid.Must(uuid.NewV4(), nil).String()
			c.Request.Header.Set(XRequestIDKey, rid)
			c.Set(XRequestIDKey, rid)
		}

		// Set XRequestIDKey header
		c.Writer.Header().Set(XRequestIDKey, rid)
		c.Next()
	}
}
