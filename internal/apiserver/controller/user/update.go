// Copyright 2021 QuanhuiLi. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package user

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	rp "github.com/liquanhui01/filestore/internal/apiserver/store/repo"
	"github.com/liquanhui01/filestore/internal/pkg/core"
)

// Update updates an user.
func (u *UserController) Update(c *gin.Context) {
	fmt.Println("user update function called.")

	var r *rp.User
	if err := c.ShouldBindJSON(&r); err != nil {
		core.WriteResponse(c, err, http.StatusInternalServerError, "服务器内部错误", nil)
	}

	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		core.WriteResponse(c, err, http.StatusBadRequest, "服务器内部错误", nil)
	}

	user, err := u.srv.Users().Find(c, uint(id))
	if err != nil {
		core.WriteResponse(c, err, http.StatusBadRequest, "服务器内部错误", nil)
	}

	user.Username = r.Username
	user.Password = r.Password
	user.Avatar = r.Avatar
	user.Email = r.Email
	user.Phone = r.Phone
	user.Profile = r.Profile

	err = u.srv.Users().Update(c, user)
	if err != nil {
		core.WriteResponse(c, err, http.StatusBadRequest, "更新失败", nil)
	}

	core.WriteResponse(c, nil, http.StatusOK, "更新成功", nil)
}
