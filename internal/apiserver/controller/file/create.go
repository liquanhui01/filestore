// Copyright 2021 QuanhuiLi. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package file

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	rp "github.com/liquanhui01/filestore/internal/apiserver/store/repo"
	"github.com/liquanhui01/filestore/internal/pkg/core"
)

func (f *FileController) Create(c *gin.Context) {
	m := make(map[string]interface{})
	var r rp.File

	if err := c.ShouldBindJSON(&r); err != nil {
		core.WriteResponse(c, err, http.StatusBadRequest, "上传失败", nil)
		return
	}

	r.Filesha1 = "dsfhdskjfheowryeoruyeiwryewire01"
	r.Location = "/Users/apple/file" + r.Filesha1
	fmt.Print(r)

	id, err := f.srv.Files().Create(c, &r)
	if err != nil {
		core.WriteResponse(c, err, http.StatusBadRequest, "上传失败", nil)
		return
	}
	m["id"] = id

	core.WriteResponse(c, nil, http.StatusCreated, "上传成功", m)
}
