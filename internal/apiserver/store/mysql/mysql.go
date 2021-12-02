package mysql

import (
	"fmt"
	"sync"

	"gorm.io/gorm"

	"github.com/liquanhui01/filestore/internal/apiserver/store"
	genericoptions "github.com/liquanhui01/filestore/internal/pkg/options"
	"github.com/liquanhui01/filestore/pkg/db"
)

type datastore struct {
	db *gorm.DB

	// include two mysql instance.
	// db *gorm.DB
	// docker *gorm.DB
}

func (ds *datastore) Users() store.UserStore {
	return newUsers(ds)
}

func (ds *datastore) Folders() store.FolderStore {
	return newFolders(ds)
}

var (
	mysqlFactory store.Factory
	once         sync.Once
)

// GetMySQLFactoryOr create mysql factory with the given config.
func GetMySQLFactoryOr(opts *genericoptions.MySQLOptions) (store.Factory, error) {
	if opts == nil && mysqlFactory == nil {
		return nil, fmt.Errorf("failed to get mysql store fatory")
	}

	var err error
	var dbIns *gorm.DB
	once.Do(func() {
		options := &db.MysQLOptions{
			Host:                  opts.Host,
			Username:              opts.Username,
			Password:              opts.Password,
			Database:              opts.Database,
			MaxIdleConnections:    opts.MaxIdleConnections,
			MaxOpenConnections:    opts.MaxOpenConnections,
			MaxConnectionLifeTime: opts.MaxConnectionLifeTime,
		}

		fmt.Println("mysql options is: ", options)

		dbIns, err = db.New(options)

		// fmt.Println("开始执行迁移")
		// err = migrateDatabase(dbIns)

		mysqlFactory = &datastore{dbIns}
	})

	if mysqlFactory == nil || err != nil {
		return nil, fmt.Errorf("failed to get mysql store factory, mysqlFactory: %+v, error: %w", mysqlFactory, err)
	}

	return mysqlFactory, nil
}

// // migrateDatabase
// func migrateDatabase(db *gorm.DB) error {
// 	if err := db.AutoMigrate(&rp.User{}); err != nil {
// 		return errors.Wrap(err, "migrate user model failed.")
// 	}

// 	return nil
// }
