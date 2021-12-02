// Copyright 2021 QuanhuiLi. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package file

import (
	srvv1 "github.com/liquanhui01/filestore/internal/apiserver/service/v1"
	"github.com/liquanhui01/filestore/internal/apiserver/store"
)

type FileController struct {
	srv srvv1.Service
}

func NewFileController(store store.Factory) *FileController {
	return &FileController{
		srv: srvv1.NewService(store),
	}
}
