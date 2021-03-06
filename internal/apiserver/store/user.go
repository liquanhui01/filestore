// Copyright 2021 QuanhuiLi. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package store

import (
	"context"

	rp "github.com/liquanhui01/filestore/internal/apiserver/store/repo"
)

type UserStore interface {
	Create(ctx context.Context, user *rp.User) (uint, error)
	Update(ctx context.Context, user *rp.User) error
	ChangePassword(ctx context.Context, id uint, password string) error
	Delete(ctx context.Context, id uint) error
	Find(ctx context.Context, id uint) (*rp.User, error)
	FindUserByName(ctx context.Context, username string) (*rp.User, error)
}
