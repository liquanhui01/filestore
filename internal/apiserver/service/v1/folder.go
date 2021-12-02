// Copyright 2021 QuanhuiLi. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package v1

import (
	"context"
	"fmt"

	"github.com/liquanhui01/filestore/internal/apiserver/store"
	rp "github.com/liquanhui01/filestore/internal/apiserver/store/repo"
)

// FolderSrv defines functions used to handle folder request.
type FolderSrv interface {
	Create(ctx context.Context, folder *rp.Folder) (uint, error)
	Update(ctx context.Context, folder *rp.Folder) error
	Find(ctx context.Context, id, userid uint) (*rp.Folder, error)
	FindByUserID(ctx context.Context, userid uint) (folders []*rp.Folder, err error)
	Delete(ctx context.Context, id, userid uint) error
}

type folderService struct {
	store store.Factory
}

var _ FolderSrv = (*folderService)(nil)

func newFolders(srv *service) *folderService {
	return &folderService{store: srv.store}
}

func (f *folderService) Create(ctx context.Context, folder *rp.Folder) (uint, error) {
	id, err := f.store.Folders().Create(ctx, folder)
	if err != nil {
		fmt.Printf("err is: %s\n", err.Error())
		return 0, err
	}

	return id, nil
}

func (f *folderService) Find(ctx context.Context, id, userid uint) (*rp.Folder, error) {
	return f.store.Folders().Find(ctx, id, userid)
}

func (f *folderService) FindByUserID(ctx context.Context, userid uint) (folders []*rp.Folder, err error) {
	return f.store.Folders().FindByUserID(ctx, userid)
}

func (f *folderService) Delete(ctx context.Context, id, userid uint) error {
	return f.store.Folders().Delete(ctx, id, userid)
}

func (f *folderService) Update(ctx context.Context, folder *rp.Folder) error {
	return f.store.Folders().Update(ctx, folder)
}
