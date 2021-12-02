// Copyright 2021 QuanhuiLi. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package file

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/liquanhui01/filestore/internal/pkg/core"
)

func (f *FileController) Find(c *gin.Context) {
	m := make(map[string]interface{})
	filesha1 := c.Param("filesha1")

	file, err := f.srv.Files().Find(c, filesha1)
	if err != nil {
		core.WriteResponse(c, err, http.StatusBadRequest, "查找失败", nil)
		return
	}
	m["file"] = file

	core.WriteResponse(c, nil, http.StatusOK, "查找成功", m)
}
