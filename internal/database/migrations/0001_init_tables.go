package migrations

import (
	"main/internal/models"

	"gorm.io/gorm"
)

type InitialMigration struct{}

func (m *InitialMigration) ID() string {
	return "0001_init_tables"
}

func (m *InitialMigration) Migrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&models.User{},
		&models.Profile{},
	)
}

func (m *InitialMigration) Rollback(db *gorm.DB) error {
	return db.Migrator().DropTable(
		&models.User{},
		&models.Profile{},
	)
}
