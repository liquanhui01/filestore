// Copyright 2021 Quanhui Li. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package middleware

// func JWTAuthMiddleware() func(c *gin.Context) {
// 	return func(c *gin.Context) {
// 		url := strings.SplitN(c.Request.URL.Path, "/", 5)
// 		if url[4] == "login" || url[4] == "logout" {
// 			c.Next()
// 			return
// 		}

// 		authHeader := c.Request.Header.Get("Authorization")
// 		if authHeader == "" {
// 			core.WriteResponse(c, nil, http.StatusBadRequest, "请求头中的auth为空", nil)
// 			c.Abort()
// 			return
// 		}

// 		parts := strings.SplitN(authHeader, " ", 2)
// 		if len(parts) == 2 {
// 			core.WriteResponse(c, nil, http.StatusBadRequest, "token不合法", nil)
// 			c.Abort()
// 			return
// 		}

// 		_, err := auth.ParseToken(parts[1])
// 		if err != nil {
// 			core.WriteResponse(c, err, http.StatusBadRequest, "token不合法", nil)
// 			c.Abort()
// 			return
// 		}

// 		c.Next()
// 	}
// }
