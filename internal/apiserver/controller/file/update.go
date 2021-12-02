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

	r.Folderid = file.Folderid
	r.Userid = file.Userid
	r.CreatedAt = file.CreatedAt
	r.Filesha1 = "ewr89r73297r382uiewuyrwruiwryu32"
	r.Filesize = 2382432
	r.Location = "/Users/apple/file" + r.Filesha1
	fmt.Println("更新后的数据：", r)

	if err := f.srv.Files().Update(c, &r); err != nil {
		core.WriteResponse(c, err, http.StatusBadRequest, "更新失败", nil)
		return
	}

	core.WriteResponse(c, nil, http.StatusOK, "更新成功", nil)
}
