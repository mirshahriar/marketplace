package migration

import (
	"fmt"

	"github.com/blang/semver"
	"github.com/mirshahriar/marketplace/helper/errors"
	"github.com/mirshahriar/marketplace/internal/ports/types"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// Migrate is the function that migrates the database to the latest schema
func Migrate(db *gorm.DB, currentVersion semver.Version) errors.Error {
	fmt.Println("Schema version is ", currentVersion.String(), ", latest version is ", latestVersion())

	for _, mi := range migrations {
		if !currentVersion.EQ(mi.fromVersion) {
			continue
		}

		err := func() errors.Error {
			fmt.Println("Migrating schema from ", currentVersion, " to ", mi.toVersion)
			tx := db.Begin()
			defer tx.Rollback()

			if err := mi.migrationFunc(tx); err != nil {
				return errors.DBMigrationError(err)
			}

			currentVersion = mi.toVersion
			system := types.SystemConfig{
				Name:  "DatabaseVersion",
				Value: currentVersion.String(),
			}

			err := db.Clauses(clause.OnConflict{UpdateAll: true}).Save(&system).Error
			if err != nil {
				return errors.DBMigrationError(err)
			}

			err = tx.Commit().Error
			if err != nil {
				return errors.DBMigrationError(err)
			}

			return nil
		}()

		if err != nil {
			return err
		}
	}
	return nil
}
