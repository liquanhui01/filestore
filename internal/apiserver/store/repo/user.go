package repo

import "time"

type User struct {
	ID        uint      `gorm:"primaryKey;autoIncrement;" json:"id"`
	Username  string    `gorm:"not null;comment:'用户名'" json:"user_name" validate:"required,min=2,max=30"`
	Password  string    `gorm:"not null;comment:'用户密码'" json:"password" validate:"required,min=6,max=20"`
	Email     string    `gorm:"default:'';comment:'邮箱'" json:"email" validate:"omitempty"`
	Phone     string    `gorm:"default:'';comment:'联系方式'" json:"phone" validate:"omitempty"`
	Avatar    string    `gorm:"default:'';comment:'用户头像'" json:"avatar" validate:"omitempty"`
	Profile   string    `gorm:"default:'';comment:'用户简介'" json:"profile" validate:"omitempty"`
	Status    int64     `gorm:"comment:'0表示启用，1表示禁用，2表示锁定，3表示标记删除'" json:"status" validate:"omitempty"`
	CreateAt  time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (u *User) TableName() string {
	return "user"
}
