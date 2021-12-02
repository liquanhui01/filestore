// Copyright 2021 QuanhuiLi. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package store

import (
	"context"

	rp "github.com/liquanhui01/filestore/internal/apiserver/store/repo"
)

type FileStore interface {
	Create(ctx context.Context, file *rp.File) (uint, error)
	Find(ctx context.Context, filesha1 string) (*rp.File, error)
	FindFilesByUserAndFolder(ctx context.Context, userid, folderid uint) (files []*rp.File, err error)
	Update(ctx context.Context, file *rp.File) error
	MoveFileToFolder(ctx context.Context, filesha1 string, folderid uint) error
	Delete(ctx context.Context, filesha1 string) error
}
