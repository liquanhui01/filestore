// Copyright 2021 QuanhuiLi. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package store

import (
	"context"

	rp "github.com/liquanhui01/filestore/internal/apiserver/store/repo"
)

type FolderStore interface {
	Create(ctx context.Context, folder *rp.Folder) (uint, error)
	Update(ctx context.Context, folder *rp.Folder) error
	Find(ctx context.Context, id, userid uint) (*rp.Folder, error)
	FindByUserID(ctx context.Context, userid uint) (folders []*rp.Folder, err error)
	Delete(ctx context.Context, id, userid uint) error
}
