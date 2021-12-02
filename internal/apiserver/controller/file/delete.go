// Copyright 2021 QuanhuiLi. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package file

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/liquanhui01/filestore/internal/pkg/core"
)

func (f *FileController) Delete(c *gin.Context) {
	filesha1 := c.Param("filesha1")

	_, err := f.srv.Files().Find(c, filesha1)
	if err != nil {
		core.WriteResponse(c, err, http.StatusNotFound, "未查找到该文件", nil)
		return
	}

	if err := f.srv.Files().Delete(c, filesha1); err != nil {
		core.WriteResponse(c, err, http.StatusBadRequest, "删除失败", nil)
		return
	}

	core.WriteResponse(c, nil, http.StatusOK, "删除成功", nil)
}
