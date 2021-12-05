// Copyright 2021 QuanhuiLi. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package file

import (
	"net/http"

	"github.com/gin-gonic/gin"

	rp "github.com/liquanhui01/filestore/internal/apiserver/store/repo"
	"github.com/liquanhui01/filestore/internal/pkg/core"
	"github.com/liquanhui01/filestore/pkg/util"
)

func (f *FileController) Update(c *gin.Context) {
	var r rp.File

	if err := c.ShouldBindJSON(&r); err != nil {
		core.WriteResponse(c, err, http.StatusBadRequest, "更新失败", nil)
		return
	}

	file, err := f.srv.Files().Find(c, r.Filesha1)
	if err != nil {
		core.WriteResponse(c, err, http.StatusNotFound, "未查找到该文件", nil)
		return
	}

	r.Filename = util.Filename(file.Filename, r.Filename)
	r.Folderid = file.Folderid
	r.Userid = file.Userid
	r.CreatedAt = file.CreatedAt
	r.Filesha1 = file.Filesha1
	r.Filesize = file.Filesize
	r.Location = file.Location

	if err := f.srv.Files().Update(c, &r); err != nil {
		core.WriteResponse(c, err, http.StatusBadRequest, "更新失败", nil)
		return
	}

	core.WriteResponse(c, nil, http.StatusOK, "更新成功", nil)
}
