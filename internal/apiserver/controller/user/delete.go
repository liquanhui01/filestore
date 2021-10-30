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

// Delete deletes an user by id
func (u *UserController) Delete(c *gin.Context) {
	fmt.Println("user delete function called.")

	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	if err := u.srv.Users().Delete(c, uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg":  "内部服务错误",
			"data": nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg":  "删除成功",
		"data": nil,
	})
}
