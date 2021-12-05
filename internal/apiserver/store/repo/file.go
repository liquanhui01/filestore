package repo

import "time"

type File struct {
	ID        uint      `gorm:"primaryKey;autoIncrement;" json:"id"`
	Filename  string    `gorm:"not null;" json:"filename"`
	Filesha1  string    `gorm:"not null;uniqueIndex:user_folder_file;" json:"filesha1"`
	Filesize  string    `gorm:"not null;" json:"filesize"`
	Location  string    `gorm:"not null;" json:"location"`
	Folderid  uint      `gorm:"not null;uniqueIndex:user_folder_file;" json:"folderid"`
	Folder    Folder    `gorm:"foreignKey:Folderid"`
	Userid    uint      `gorm:"not null;uniqueIndex:user_folder_file;" json:"userid"`
	User      User      `gorm:"foreignKey:Userid"`
	CreatedAt time.Time `gorm:"autoCreateTime:nano;" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime:nano;" json:"updated_at"`
}

func (f *File) TableName() string {
	return "file"
}
