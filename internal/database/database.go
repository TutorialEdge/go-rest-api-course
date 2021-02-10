package database

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// NewDatabase - returns a pointer to a new database connection
func NewDatabase() (*gorm.DB, error) {
	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbTable := os.Getenv("DB_TABLE")
	dbPort := os.Getenv("DB_PORT")

	// dbConnectionString := dbUsername + ":" + dbPassword + "@tcp(" + dbHost + ":" + strconv.Itoa(dbPort) + ")/" + dbTable
	connectionString := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", dbHost, dbPort, dbUsername, dbTable, dbPassword)
	fmt.Println(connectionString)

	db, err := gorm.Open("postgres", connectionString)
	if err != nil {
		return db, err
	}

	if err := db.DB().Ping(); err != nil {
		return db, err
	}

	return db, nil
}
