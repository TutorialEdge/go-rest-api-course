package database

import "github.com/jinzhu/gorm"

// NewDatabase - returns a pointer to a new database connection
func NewDatabase() (*gorm.DB, error) {
	db, err := gorm.Open()
	if err != nil {
		return db, err
	}
	return db, nil
}
