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

// UserSrv defines functions used to handle user request.
type UserSrv interface {
	Create(ctx context.Context, user *rp.User) (uint, error)
	Update(ctx context.Context, user *rp.User) error
	ChangePassword(ctx context.Context, id uint, password string) error
	Delete(ctx context.Context, id uint) error
	Find(ctx context.Context, id uint) (*rp.User, error)
}

type userService struct {
	store store.Factory
}

var _ UserSrv = (*userService)(nil)

func newUsers(srv *service) *userService {
	return &userService{store: srv.store}
}

func (u *userService) Create(ctx context.Context, user *rp.User) (uint, error) {
	id, err := u.store.Users().Create(ctx, user)
	if err != nil {
		fmt.Printf("err is:%s\n", err.Error())
		return 0, err
	}

	return id, nil
}

func (u *userService) Delete(ctx context.Context, id uint) error {
	return u.store.Users().Delete(ctx, id)
}

func (u *userService) Update(ctx context.Context, user *rp.User) error {
	return u.store.Users().Update(ctx, user)
}

func (u *userService) ChangePassword(ctx context.Context, id uint, password string) error {
	return u.store.Users().ChangePassword(ctx, id, password)
}

func (u *userService) Find(ctx context.Context, id uint) (user *rp.User, err error) {
	user, err = u.store.Users().Find(ctx, id)
	if err != nil {
		return nil, err
	}
	return user, nil
}
