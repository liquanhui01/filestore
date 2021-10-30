// Copyright 2021 Quanhui Li. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package middleware

import "github.com/gin-gonic/gin"

const UsernameKey = "username"

// Context is a middleware that injects common prefix fields to gin.Context.
func Context() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("KeyRequestID", c.GetString(XRequestIDKey))
		c.Set("KeyUsername", c.GetString(UsernameKey))
		c.Next()
	}
}
