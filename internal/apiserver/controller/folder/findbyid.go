// Copyright 2021 QuanhuiLi. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package folder

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/liquanhui01/filestore/internal/pkg/core"
)

func (f *FolderController) FindByUserID(c *gin.Context) {
	m := make(map[string]interface{})
	userid, _ := strconv.ParseUint(c.Param("userid"), 10, 64)

	folders, err := f.srv.Folders().FindByUserID(c, uint(userid))
	if err != nil {
		core.WriteResponse(c, err, http.StatusInternalServerError, "查找失败", nil)
		return
	}
	m["folders"] = folders

	core.WriteResponse(c, nil, http.StatusOK, "查找成功", m)
}
