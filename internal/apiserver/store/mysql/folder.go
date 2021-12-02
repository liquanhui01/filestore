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

type folders struct {
	db *gorm.DB
}

func newFolders(ds *datastore) *folders {
	return &folders{db: ds.db}
}

func (f *folders) Create(ctx context.Context, folder *rp.Folder) (uint, error) {
	err := f.db.WithContext(ctx).Create(&folder).Error
	if err != nil {
		fmt.Printf("错误： %s\n", err.Error())
	}

	return folder.ID, err
}

// Update update folder
func (f *folders) Update(ctx context.Context, folder *rp.Folder) error {
	return f.db.WithContext(ctx).Model(&folder).Updates(folder).Error
}

// Find folder by userid and id
func (f *folders) Find(ctx context.Context, id, userid uint) (folder *rp.Folder, err error) {
	return folder, f.db.WithContext(ctx).Where("id = ? And user_id = ?", id, userid).First(&folder).Error
}

// FindByUserID find folders by userid
func (f *folders) FindByUserID(ctx context.Context, userid uint) (folders []*rp.Folder, err error) {
	return folders, f.db.WithContext(ctx).Where("user_id = ?", userid).Find(&folders).Error
}

// delete folder by id
func (f *folders) Delete(ctx context.Context, id, userid uint) error {
	return f.db.WithContext(ctx).Where("id = ? And user_id = ?", id, userid).Delete(&rp.Folder{}).Error
}
