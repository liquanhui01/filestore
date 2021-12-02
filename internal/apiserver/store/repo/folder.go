package repo

import "time"

type Folder struct {
	ID         uint      `gorm:"primaryKey;autoIncrement;" json:"id"`
	Foldername string    `gorm:"not null; comment:'文件夹名称'" json:"foldername"`
	UserID     uint      `gorm:"not null; comment:'文件夹所属用户'" json:"userid"`
	User       User      `gorm:"->:false;<-:create"`
	CreatedAt  time.Time `gorm:"autoCreateTime:nano;" json:"created_at"`
	UpdatedAt  time.Time `gorm:"autoUpdateTime:nano;" json:"updated_at"`
}

func (f *Folder) TableName() string {
	return "folder"
}
