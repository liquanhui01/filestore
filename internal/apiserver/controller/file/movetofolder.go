// Copyright 2021 QuanhuiLi. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package file

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/liquanhui01/filestore/internal/pkg/core"
)

type FileInfo struct {
	Filesha1 string `json:"filesha1"`
	UserID   uint   `json:"userid"`
	FolderID uint   `json:"folderid"`
}

func (f *FileController) MoveFileToFolder(c *gin.Context) {
	var r FileInfo

	if err := c.ShouldBindJSON(&r); err != nil {
		core.WriteResponse(c, err, http.StatusBadRequest, "移动文件失败", nil)
		return
	}

	if err := f.srv.Files().MoveFileToFolder(c, r.Filesha1, r.FolderID); err != nil {
		core.WriteResponse(c, err, http.StatusBadRequest, "移动文件失败", nil)
		return
	}

	core.WriteResponse(c, nil, http.StatusOK, "移动成功", nil)
}
