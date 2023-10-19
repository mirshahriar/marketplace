package migration

import (
	"github.com/blang/semver"
	"github.com/mirshahriar/marketplace/internal/framework/mysql/migration/queries"
	"gorm.io/gorm"
)

type migration struct {
	fromVersion   semver.Version
	toVersion     semver.Version
	migrationFunc func(tx *gorm.DB) error
}

func latestVersion() semver.Version {
	return migrations[len(migrations)-1].toVersion
}

// migrations hold all the schema migrations
var migrations = []migration{
	{
		semver.MustParse("0.0.0"),
		semver.MustParse("0.1.0"),
		queries.GetMigrationQueryV010(),
	},
}
