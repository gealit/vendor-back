package migrations

import (
	"fmt"

	"gorm.io/gorm"
)

type Migration interface {
	ID() string
	Migrate(*gorm.DB) error
	Rollback(*gorm.DB) error
}

var AllMigrations = []Migration{
	&InitialMigration{},
}

func RunMigrations(db *gorm.DB) error {
	for _, migration := range AllMigrations {
		if err := migration.Migrate(db); err != nil {
			return fmt.Errorf("failed to run migration %s: %w", migration.ID(), err)
		}
	}
	return nil
}
