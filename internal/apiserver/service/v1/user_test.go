// Copyright 2021 QuanhuiLi. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package v1

import (
	"context"
	"reflect"
	"testing"

	"github.com/liquanhui01/filestore/internal/apiserver/store"

	gomock "github.com/golang/mock/gomock"
	rp "github.com/liquanhui01/filestore/internal/apiserver/store/repo"
	"github.com/stretchr/testify/suite"
)

type Suite struct {
	suite.Suite
	MockFactory *store.MockFactory

	mockUserStore *store.MockUserStore
	users         []*rp.User
}

func Test_newUsers(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockFactory := store.NewMockFactory(ctrl)

	type args struct {
		srv *service
	}
	tests := []struct {
		name string
		args args
		want *userService
	}{
		// TODO: Add test cases.
		{
			name: "default",
			args: args{
				srv: &service{store: mockFactory},
			},
			want: &userService{
				store: mockFactory,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newUsers(tt.args.srv); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newUsers() = %v, want %v", got, tt.want)
			}
		})
	}
}

func (s *Suite) Test_userService_Create(t *testing.T) {
	s.mockUserStore.EXPECT().Create(gomock.Any(), gomock.Eq(s.users[0])).Times(1000000).Return(nil)
	type fields struct {
		store store.Factory
	}
	type args struct {
		ctx  context.Context
		user *rp.User
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    uint
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "default",
			fields: fields{
				store: s.MockFactory,
			},
			args: args{
				ctx:  context.TODO(),
				user: s.users[0],
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &userService{
				store: tt.fields.store,
			}
			got, err := u.Create(tt.args.ctx, tt.args.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("userService.Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("userService.Create() = %v, want %v", got, tt.want)
			}
		})
	}
}

func (s *Suite) Test_userService_Delete(t *testing.T) {
	s.mockUserStore.EXPECT().Delete(gomock.Any(), 1).Return().Times(100000)
	type fields struct {
		store store.Factory
	}
	type args struct {
		ctx context.Context
		id  uint
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "default",
			fields: fields{
				store: s.MockFactory,
			},
			args: args{
				ctx: context.TODO(),
				id:  1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &userService{
				store: tt.fields.store,
			}
			if err := u.Delete(tt.args.ctx, tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("userService.Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func (s *Suite) Test_userService_Update(t *testing.T) {
	type fields struct {
		store store.Factory
	}
	type args struct {
		ctx  context.Context
		user *rp.User
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &userService{
				store: tt.fields.store,
			}
			if err := u.Update(tt.args.ctx, tt.args.user); (err != nil) != tt.wantErr {
				t.Errorf("userService.Update() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func (s *Suite) Test_userService_Find(t *testing.T) {
	type fields struct {
		store store.Factory
	}
	type args struct {
		ctx context.Context
		id  uint
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		wantUser *rp.User
		wantErr  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &userService{
				store: tt.fields.store,
			}
			gotUser, err := u.Find(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("userService.Find() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotUser, tt.wantUser) {
				t.Errorf("userService.Find() = %v, want %v", gotUser, tt.wantUser)
			}
		})
	}
}
