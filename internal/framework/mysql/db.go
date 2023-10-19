// Package db implements DBPort with mysql
package db

import (
	"fmt"
	"time"

	"github.com/mirshahriar/marketplace/config"
	"github.com/mirshahriar/marketplace/helper/errors"
	"github.com/mirshahriar/marketplace/internal/ports"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Adapter implements ports.DBPort with mysql
type Adapter struct {
	db *gorm.DB
}

// This validates that Adapter implements the ports.DBPort interface
var _ ports.DBPort = Adapter{}

// NewAdapterWithConfig creates a new adapter with the given config
func NewAdapterWithConfig(cfg config.DBConfig) (*Adapter, errors.Error) {
	var adapter *Adapter
	var cErr errors.Error

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Name)
	dialect := mysql.Open(dsn)
	adapter, cErr = newAdapterWithDialector(dialect)
	if cErr != nil {
		return nil, cErr
	}

	sqlDB, err := adapter.db.DB()
	if err != nil {
		return nil, errors.InternalError(err)
	}

	sqlDB.SetConnMaxLifetime(cfg.MaxConnTime)
	sqlDB.SetMaxIdleConns(cfg.MaxIdleConn)
	sqlDB.SetMaxOpenConns(cfg.MaxOpenConn)

	return adapter, nil
}

// newAdapterWithDialector creates a new adapter with the given dialect
func newAdapterWithDialector(dialect gorm.Dialector) (*Adapter, errors.Error) {
	d, err := gorm.Open(dialect, &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
		NowFunc: func() time.Time {
			return time.Now()
		},
	})
	if err != nil {
		return nil, errors.InternalError(err)
	}

	adapter := &Adapter{
		db: d,
	}

	//if cErr := adapter.migrate(); cErr != nil {
	//	return nil, cErr
	//}

	createCallback := d.Callback().Create()
	createCallback.Clauses = []string{"INSERT", "VALUES"}

	return adapter, nil
}

// nolint: unused
func (a Adapter) migrate() errors.Error {
	db := a.db.Session(&gorm.Session{
		Logger: logger.Default.LogMode(logger.Error)})

	tables := []interface{}{}

	for _, table := range tables {
		if !db.Migrator().HasTable(table) {
			err := db.Migrator().CreateTable(table)
			if err != nil {
				return errors.InternalDBError(err)
			}
		}
	}

	return nil
}
