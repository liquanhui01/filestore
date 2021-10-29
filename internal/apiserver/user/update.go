// Copyright 2021 QuanhuiLi. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package user

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	rp "file-store/internal/store/repo"
)

// Update updates an user.
func (u *UserController) Update(c *gin.Context) {
	fmt.Println("user update function called.")

	var r *rp.User

	if err := c.ShouldBindJSON(&r); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg":  "服务器内部错误",
			"data": nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg":  "更新成功",
		"data": nil,
	})
}
