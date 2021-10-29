// Copyright 2021 QuanhuiLi. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package user

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Find get an user by id
func (u *UserController) Find(c *gin.Context) {
	fmt.Println("find user function called.")

	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	user, err := u.srv.Users().Find(c, uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg":  "服务器错误",
			"data": nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg":  "查询成功",
		"data": user,
	})
}
