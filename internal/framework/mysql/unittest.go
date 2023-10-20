package db

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/mirshahriar/marketplace/internal/ports/types"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func NewAdapterForTest(t *testing.T) (*Adapter, error) {

	name := fmt.Sprintf("gorm_%s.db", strings.ToLower(t.Name()))

	// Cleanup the sqlite file after the test
	t.Cleanup(func() {
		_ = os.Remove(name)
	})

	db, err := gorm.Open(sqlite.Open(name), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	adapter := &Adapter{
		db: db,
	}

	// Enable parallel testing
	t.Parallel()

	// Migrate the schema
	err = db.AutoMigrate(&types.Product{})
	if err != nil {
		return nil, err
	}

	return adapter, nil
}
