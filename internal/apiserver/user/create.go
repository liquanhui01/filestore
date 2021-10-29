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

// Create add a new user to mysql.
func (u *UserController) Create(c *gin.Context) {
	fmt.Println("user create function called.")

	var r rp.User

	if err := c.ShouldBindJSON(&r); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg":  "服务器内部错误",
			"data": nil,
		})
		return
	}

	// Insert the user to mysql.
	id, err := u.srv.Users().Create(c, &r)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg":  "服务器内部错误",
			"data": nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg":  "注册成功",
		"data": id,
	})
}
