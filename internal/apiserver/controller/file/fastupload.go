// Copyright 2021 QuanhuiLi. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package file

import (
	"net/http"

	"github.com/gin-gonic/gin"

	rp "github.com/liquanhui01/filestore/internal/apiserver/store/repo"
	"github.com/liquanhui01/filestore/internal/pkg/core"
)

func (f *FileController) FastUpload(c *gin.Context) {
	var r rp.File
	m := make(map[string]interface{})

	if err := c.ShouldBindJSON(&r); err != nil {
		core.WriteResponse(c, err, http.StatusInternalServerError, "秒传失败，请重试", nil)
		return
	}

	userfile, err := f.srv.Files().Find(c, r.Filesha1)
	if err != nil {
		core.WriteResponse(c, err, http.StatusBadRequest, "秒传失败，请使用普通上传", nil)
		return
	}

	r.Location = userfile.Location
	r.Filesize = userfile.Filesize

	id, err := f.srv.Files().Create(c, &r)
	if err != nil {
		core.WriteResponse(c, err, http.StatusInternalServerError, "秒传失败，请重试", nil)
		return
	}
	m["id"] = id

	core.WriteResponse(c, nil, http.StatusCreated, "秒传成功", m)
}
