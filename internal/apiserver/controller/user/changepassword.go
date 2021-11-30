// Copyright 2021 QuanhuiLi. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package user

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/liquanhui01/filestore/internal/pkg/core"
)

type Password struct {
	OldPassword string `json:"old_password"`
	NewPassword string `json:"new_password"`
}

func (u *UserController) ChangePassword(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		core.WriteResponse(c, err, http.StatusInternalServerError, "修改密码失败", nil)
		return
	}

	var r Password

	if err = c.ShouldBindJSON(&r); err != nil {
		core.WriteResponse(c, err, http.StatusInternalServerError, "修改密码失败", nil)
		return
	}

	user, err := u.srv.Users().Find(c, uint(id))
	if err != nil {
		core.WriteResponse(c, err, http.StatusBadRequest, "用户不存在", nil)
		return
	}

	if err = core.ValidatePassword(r.OldPassword, user.Password); err != nil {
		core.WriteResponse(c, err, http.StatusBadRequest, "密码不匹配", nil)
		return
	}

	password, err := core.GeneratePassword(r.NewPassword)
	if err != nil {
		core.WriteResponse(c, err, http.StatusInternalServerError, "修改密码失败", nil)
		return
	}
	newPassword := string(password)

	if err = u.srv.Users().ChangePassword(c, uint(id), newPassword); err != nil {
		core.WriteResponse(c, err, http.StatusBadRequest, "修改密码失败", nil)
		return
	}

	core.WriteResponse(c, nil, http.StatusOK, "修改密码成功", nil)
}
