// Copyright 2021 QuanhuiLi. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package folder

import (
	"net/http"

	"github.com/gin-gonic/gin"

	rp "github.com/liquanhui01/filestore/internal/apiserver/store/repo"
	"github.com/liquanhui01/filestore/internal/pkg/core"
)

func (f *FolderController) Create(c *gin.Context) {
	var r rp.Folder
	mp := make(map[string]interface{})

	if err := c.ShouldBindJSON(&r); err != nil {
		core.WriteResponse(c, err, http.StatusInternalServerError, "创建失败", nil)
		return
	}

	id, err := f.srv.Folders().Create(c, &r)
	if err != nil {
		core.WriteResponse(c, err, http.StatusBadRequest, "创建失败", nil)
		return
	}
	mp["id"] = id

	core.WriteResponse(c, nil, http.StatusCreated, "创建成功", mp)
}
