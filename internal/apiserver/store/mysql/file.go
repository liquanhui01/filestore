// Copyright 2021 QuanhuiLi. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package mysql

import (
	"context"
	"fmt"

	"gorm.io/gorm"

	rp "github.com/liquanhui01/filestore/internal/apiserver/store/repo"
)

type file struct {
	db *gorm.DB
}

func newFile(ds *datastore) *file {
	return &file{db: ds.db}
}

// Create create files.
func (f *file) Create(ctx context.Context, file *rp.File) (uint, error) {
	err := f.db.WithContext(ctx).Create(&file).Error
	if err != nil {
		fmt.Printf("err is： %s\n", err.Error())
	}

	return file.ID, err
}

// Update update file.
func (f *file) Update(ctx context.Context, file *rp.File) error {
	return f.db.WithContext(ctx).Model(&file).Where("filesha1 = ?", file.Filesha1).Select("*").Omit("folderid").Save(&file).Error
}

// MoveFileToFolder
func (f *file) MoveFileToFolder(ctx context.Context, filesha1 string, folderid uint) error {
	return f.db.WithContext(ctx).Model(&rp.File{}).Where("filesha1 = ?", filesha1).Update("folderid", folderid).Error
}

// Find find file by filesha1.
func (f *file) Find(ctx context.Context, filesha1 string) (file *rp.File, err error) {
	return file, f.db.WithContext(ctx).Where("filesha1 = ?", filesha1).First(&file).Error
}

// FindFilesByUserAndFolder
func (f *file) FindFilesByUserAndFolder(ctx context.Context, userid, folderid uint) (files []*rp.File, err error) {
	return files, f.db.WithContext(ctx).Where("userid = ? And folderid = ?", userid, folderid).Find(&files).Error
}

// Delete delete file by filesha1.
func (f *file) Delete(ctx context.Context, filesha1 string) error {
	return f.db.WithContext(ctx).Where("filesha1 = ?", filesha1).Delete(&rp.File{}).Error
}
