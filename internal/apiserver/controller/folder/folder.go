// Copyright 2021 QuanhuiLi. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package folder

import (
	srvv1 "github.com/liquanhui01/filestore/internal/apiserver/service/v1"
	"github.com/liquanhui01/filestore/internal/apiserver/store"
)

type FolderController struct {
	srv srvv1.Service
}

func NewFolderController(store store.Factory) *FolderController {
	return &FolderController{
		srv: srvv1.NewService(store),
	}
}
