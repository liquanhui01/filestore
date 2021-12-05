package mysql

import (
	"context"
	"fmt"

	"gorm.io/gorm"

	rp "github.com/liquanhui01/filestore/internal/apiserver/store/repo"
)

type users struct {
	db *gorm.DB
}

func newUsers(ds *datastore) *users {
	return &users{db: ds.db}
}

// Create creates new user, return id and nil if err is nil.
func (u *users) Create(ctx context.Context, user *rp.User) (uint, error) {
	err := u.db.WithContext(ctx).Create(&user).Error
	if err != nil {
		fmt.Printf("错误：%s\n", err.Error())
	}

	return user.ID, err
}

// Update updates user
func (u *users) Update(ctx context.Context, user *rp.User) error {
	return u.db.WithContext(ctx).Model(&rp.User{}).Where("id = ?", user.ID).Select("*").Omit("Password").Save(&user).Error
}

// Update user's password
func (u *users) ChangePassword(ctx context.Context, id uint, password string) error {
	return u.db.WithContext(ctx).Model(&rp.User{}).Where("id = ?", id).Update("password", password).Error
}

// Delete deletes user by id
func (u *users) Delete(ctx context.Context, id uint) error {
	return u.db.WithContext(ctx).Delete(&rp.User{}, id).Error
}

// Find find user by id
func (u *users) Find(ctx context.Context, id uint) (user *rp.User, err error) {
	return user, u.db.WithContext(ctx).First(&user, id).Error
}

// FindUserByName find user by username
func (u *users) FindUserByName(ctx context.Context, username string) (user *rp.User, err error) {
	return user, u.db.WithContext(ctx).Where("username = ?", username).First(&user).Error
}
