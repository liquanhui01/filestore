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

// Delete deletes an user by id
func (u *UserController) Delete(c *gin.Context) {
	fmt.Println("user delete function called.")

	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		core.WriteResponse(c, err, http.StatusInternalServerError, "删除失败，服务器内部错误", nil)
	}

	if err := u.srv.Users().Delete(c, uint(id)); err != nil {
		core.WriteResponse(c, err, http.StatusInternalServerError, "删除失败", nil)
	}

	core.WriteResponse(c, nil, http.StatusOK, "删除成功", nil)
}
