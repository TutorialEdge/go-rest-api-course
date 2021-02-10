package database

import "github.com/jinzhu/gorm"

// MigrateDB - migrates our db based on the models passed in
func MigrateDB(db *gorm.DB, model interface{}) {
	db.AutoMigrate(&model)
}
