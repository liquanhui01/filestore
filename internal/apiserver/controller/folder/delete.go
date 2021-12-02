// Copyright 2021 QuanhuiLi. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package folder

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/liquanhui01/filestore/internal/pkg/core"
)

type folderInfo struct {
	ID     uint `json:"id"`
	UserID uint `json:"userid"`
}

func (f *FolderController) Delete(c *gin.Context) {
	var r folderInfo
	if err := c.ShouldBindJSON(&r); err != nil {
		core.WriteResponse(c, err, http.StatusInternalServerError, "删除失败", nil)
		return
	}

	_, err := f.srv.Folders().Find(c, r.ID, r.UserID)
	if err != nil {
		core.WriteResponse(c, err, http.StatusNotFound, "该文件夹不存在", nil)
		return
	}

	if err := f.srv.Folders().Delete(c, r.ID, r.UserID); err != nil {
		core.WriteResponse(c, err, http.StatusInternalServerError, "删除失败", nil)
		return
	}

	core.WriteResponse(c, nil, http.StatusOK, "删除成功", nil)
}
