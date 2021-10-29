package mysql

import (
	"gorm.io/gorm"

	"file-store/internal/store"
)

type datastore struct {
	db *gorm.DB

	// 可以包括两个实例
	// db *gorm.DB
	// docker *gorm.DB
}

func (ds *datastore) Users() store.UserStore {
	return newUsers(ds)
}
