package models

import "gorm.io/gorm"

func MigrateModels(db *gorm.DB) error {
	return db.AutoMigrate(
		&User{},
	)
}
