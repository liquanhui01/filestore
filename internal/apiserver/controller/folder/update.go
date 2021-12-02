// Copyright 2021 Quanhui Li. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package folder

import (
	"net/http"

	"github.com/gin-gonic/gin"

	rp "github.com/liquanhui01/filestore/internal/apiserver/store/repo"
	"github.com/liquanhui01/filestore/internal/pkg/core"
)

func (f *FolderController) Update(c *gin.Context) {
	var r rp.Folder
	// id, _ := strconv.ParseUint(c.Param("id"), 10, 64)

	if err := c.ShouldBindJSON(&r); err != nil {
		core.WriteResponse(c, err, http.StatusInternalServerError, "更新失败", nil)
		return
	}

	folder, err := f.srv.Folders().Find(c, r.ID, r.UserID)
	if err != nil {
		core.WriteResponse(c, err, http.StatusNotFound, "该文件夹不存在", nil)
		return
	}
	folder.Foldername = r.Foldername

	if err := f.srv.Folders().Update(c, folder); err != nil {
		core.WriteResponse(c, err, http.StatusInternalServerError, "更新失败", nil)
		return
	}

	core.WriteResponse(c, nil, http.StatusOK, "更新成功", nil)
}
