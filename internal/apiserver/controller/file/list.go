// Copyright 2021 QuanhuiLi. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package file

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/liquanhui01/filestore/internal/pkg/core"
)

func (f *FileController) FindFilesByUserAndFolder(c *gin.Context) {
	m := make(map[string]interface{})
	var r FileInfo

	if err := c.ShouldBindJSON(&r); err != nil {
		core.WriteResponse(c, err, http.StatusInternalServerError, "查找失败", nil)
		return
	}

	files, err := f.srv.Files().FindFilesByUserAndFolder(c, r.UserID, r.FolderID)
	if err != nil {
		core.WriteResponse(c, err, http.StatusBadRequest, "查找失败", nil)
		return
	}
	m["files"] = files

	core.WriteResponse(c, nil, http.StatusOK, "查找成功", m)
}
