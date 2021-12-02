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

// FileSrv defines functions used to handle file request.
type FileSrv interface {
	Create(ctx context.Context, file *rp.File) (uint, error)
	Find(ctx context.Context, filesha1 string) (*rp.File, error)
	FindFilesByUserAndFolder(ctx context.Context, userid, folderid uint) (files []*rp.File, err error)
	Update(ctx context.Context, file *rp.File) error
	MoveFileToFolder(ctx context.Context, filesha1 string, folderid uint) error
	Delete(ctx context.Context, filesha1 string) error
}

type fileService struct {
	store store.Factory
}

var _ FileSrv = (*fileService)(nil)

func newFiles(srv *service) *fileService {
	return &fileService{store: srv.store}
}

func (f *fileService) Create(ctx context.Context, file *rp.File) (uint, error) {
	id, err := f.store.Files().Create(ctx, file)
	if err != nil {
		fmt.Printf("err is: %s\n", err.Error())
		return 0, err
	}

	return id, nil
}

func (f *fileService) Find(ctx context.Context, filesha1 string) (*rp.File, error) {
	return f.store.Files().Find(ctx, filesha1)
}

func (f *fileService) FindFilesByUserAndFolder(ctx context.Context, userid, folderid uint) (files []*rp.File, err error) {
	return f.store.Files().FindFilesByUserAndFolder(ctx, userid, folderid)
}

func (f *fileService) Update(ctx context.Context, file *rp.File) error {
	return f.store.Files().Update(ctx, file)
}

func (f *fileService) MoveFileToFolder(ctx context.Context, filesha1 string, folderid uint) error {
	return f.store.Files().MoveFileToFolder(ctx, filesha1, folderid)
}

func (f *fileService) Delete(ctx context.Context, filesha1 string) error {
	return f.store.Files().Delete(ctx, filesha1)
}
