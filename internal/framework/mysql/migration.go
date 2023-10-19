package db

import (
	"github.com/blang/semver"
	"github.com/mirshahriar/marketplace/helper/errors"
	"github.com/mirshahriar/marketplace/internal/framework/mysql/migration"
	"github.com/mirshahriar/marketplace/internal/ports/types"
)

// Migrate migrates the database to the latest version
func (a Adapter) Migrate() errors.Error {
	err := a.db.AutoMigrate(&types.SystemConfig{})
	if err != nil {
		return errors.InternalDBError(err)
	}

	var system types.SystemConfig
	rowAffected := a.db.Model(&types.SystemConfig{}).
		Where(&types.SystemConfig{Name: "DatabaseVersion"}).
		Find(&system).RowsAffected

	if rowAffected == 0 {
		system.Value = "0.0.0"
	}

	currentVersion, err := semver.Parse(system.Value)
	if err != nil {
		return errors.DBMigrationError(err)
	}

	return migration.Migrate(a.db, currentVersion)
}
