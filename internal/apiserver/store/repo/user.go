package repo

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"not null;unique;comment:'用户名'" json:"user_name" validate:"required,min=2,max=30"`
	Password string `gorm:"not null;comment:'用户密码'" json:"password" validate:"required,min=6,max=20"`
	Email    string `gorm:"default:'';unique;comment:'邮箱'" json:"email" validate:"omitempty"`
	Phone    string `gorm:"default:'';unique;comment:'联系方式'" json:"phone" validate:"omitempty"`
	Avatar   string `gorm:"default:'';comment:'用户头像'" json:"avatar" validate:"omitempty"`
	Profile  string `gorm:"default:'';comment:'用户简介'" json:"profile" validate:"omitempty"`
	Status   int64  `gorm:"comment:'0表示启用，1表示禁用，2表示锁定，3表示标记删除'" json:"status" validate:"omitempty"`
}

func (u *User) TableName() string {
	return "user"
}
