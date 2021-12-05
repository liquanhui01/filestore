// Copyright 2021 QuanhuiLi. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package file

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	rp "github.com/liquanhui01/filestore/internal/apiserver/store/repo"
	"github.com/liquanhui01/filestore/internal/pkg/core"
	"github.com/liquanhui01/filestore/pkg/util"
)

func (f *FileController) Upload(c *gin.Context) {
	m := make(map[string]interface{})
	var r rp.File

	userid, _ := strconv.ParseUint(c.Param("userid"), 10, 64)
	folderid, _ := strconv.ParseUint(c.Param("folderid"), 10, 64)

	if err := c.Request.ParseMultipartForm(4 << 20); err != nil {
		core.WriteResponse(c, err, http.StatusBadRequest, "上传大小超过限制", nil)
		return
	}

	file, head, err := c.Request.FormFile("file")
	if err != nil {
		core.WriteResponse(c, err, http.StatusBadRequest, "上传失败", nil)
		return
	}
	defer file.Close()

	r, err = util.Meta(r, folderid, userid, head.Filename, file)
	if err != nil {
		core.WriteResponse(c, err, http.StatusInternalServerError, "上传失败", nil)
		return
	}

	id, err := f.srv.Files().Create(c, &r)
	if err != nil {
		core.WriteResponse(c, err, http.StatusBadRequest, "上传失败", nil)
		return
	}
	m["id"] = id

	core.WriteResponse(c, nil, http.StatusCreated, "上传成功", m)
}
