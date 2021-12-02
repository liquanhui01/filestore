// Copyright 2021 QuanhuiLi. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package folder

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/liquanhui01/filestore/internal/pkg/core"
)

func (f *FolderController) Find(c *gin.Context) {
	m := make(map[string]interface{})
	var r folderInfo

	if err := c.ShouldBindJSON(&r); err != nil {
		core.WriteResponse(c, err, http.StatusInternalServerError, "查找失败", nil)
		return
	}

	folder, err := f.srv.Folders().Find(c, r.ID, r.UserID)
	if err != nil {
		core.WriteResponse(c, err, http.StatusInternalServerError, "查找失败", nil)
		return
	}
	m["folder"] = folder

	core.WriteResponse(c, nil, http.StatusOK, "查找成功", m)
}
