// Copyright 2021 QuanhuiLi. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package user

import (
	"net/http"

	"github.com/gin-gonic/gin"

	rp "github.com/liquanhui01/filestore/internal/apiserver/store/repo"
	"github.com/liquanhui01/filestore/internal/pkg/core"
)

// Create add a new user to mysql.
func (u *UserController) Create(c *gin.Context) {
	var r rp.User
	mp := make(map[string]interface{})
	if err := c.ShouldBindJSON(&r); err != nil {
		core.WriteResponse(c, err, http.StatusInternalServerError, "服务器内部错误", nil)
		return
	}

	hashed, err := core.GeneratePassword(r.Password)
	if err != nil {
		core.WriteResponse(c, err, http.StatusInternalServerError, "服务器内部错误", nil)
		return
	}
	r.Password = string(hashed)

	// Insert the user to mysql.
	id, err := u.srv.Users().Create(c, &r)
	if err != nil {
		core.WriteResponse(c, err, http.StatusInternalServerError, "用户创建失败", nil)
		return
	}
	mp["id"] = id

	core.WriteResponse(c, nil, http.StatusCreated, "用户创建成功", mp)
}
