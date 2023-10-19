package queries

import "gorm.io/gorm"

func GetMigrationQueryV010() func(tx *gorm.DB) error {
	return func(tx *gorm.DB) error {
		return nil
	}
}
