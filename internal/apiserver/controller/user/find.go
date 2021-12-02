// Copyright 2021 QuanhuiLi. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package user

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/liquanhui01/filestore/internal/pkg/core"
)

// Find get an user by id
func (u *UserController) Find(c *gin.Context) {
	fmt.Println("find user function called.")
	mp := make(map[string]interface{})

	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	user, err := u.srv.Users().Find(c, uint(id))
	if err != nil {
		core.WriteResponse(c, err, http.StatusInternalServerError, "未查询到当前用户", nil)
		return
	}
	mp["user"] = user

	core.WriteResponse(c, nil, http.StatusOK, "查询成功", mp)
}
