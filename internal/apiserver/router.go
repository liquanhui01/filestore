// Copyright 2021 Quanhui Li. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package apiserver

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func initRouter() *gin.Engine {
	r := gin.New()

	apiv1 := r.Group("/api/v1")
	{
		apiv1.GET("", func(c *gin.Context) {
			fmt.Println("成功")
		})
	}

	return r
}
