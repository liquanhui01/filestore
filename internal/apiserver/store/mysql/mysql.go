package mysql

import (
	"fmt"
	"sync"

	"github.com/pkg/errors"
	"gorm.io/gorm"

	"filestore/internal/apiserver/store"
	rp "filestore/internal/apiserver/store/repo"
	genericoptions "filestore/internal/pkg/options"
	"filestore/pkg/db"
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

		dbIns, err = db.New(options)

		fmt.Println("开始执行迁移")
		err = migrateDatabase(dbIns)

		mysqlFactory = &datastore{dbIns}
	})

	if mysqlFactory == nil || err != nil {
		return nil, fmt.Errorf("failed to get mysql store factory, mysqlFactory: %+v, error: %w", mysqlFactory, err)
	}

	return mysqlFactory, nil
}

// migrateDatabase
func migrateDatabase(db *gorm.DB) error {
	if err := db.AutoMigrate(&rp.User{}); err != nil {
		return errors.Wrap(err, "migrate user model failed.")
	}

	return nil
}
