package mysql

import (
	"context"

	"gorm.io/gorm"

	rp "filestore/internal/apiserver/store/repo"
)

type users struct {
	db *gorm.DB
}

func newUsers(ds *datastore) *users {
	return &users{db: ds.db}
}

// Create creates new user, return id and nil if err is nil.
func (u *users) Create(ctx context.Context, user *rp.User) (uint, error) {
	return user.ID, u.db.WithContext(ctx).Create(&user).Error
}

// Update updates user
func (u *users) Update(ctx context.Context, user *rp.User) error {
	return u.db.WithContext(ctx).Save(user).Error
}

// Delete deletes user by id
func (u *users) Delete(ctx context.Context, id uint) error {
	return u.db.WithContext(ctx).Delete(&rp.User{}, id).Error
}

// Find find user by id
func (u *users) Find(ctx context.Context, id uint) (user *rp.User, err error) {
	return user, u.db.WithContext(ctx).First(&user, id).Error
}
