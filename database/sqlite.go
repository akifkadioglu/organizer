package database

import (
	"os"
	"strings"

	"github.com/akifkadioglu/organizer/config"
	"github.com/akifkadioglu/organizer/database/entities"
	"github.com/akifkadioglu/organizer/olog"
	"gorm.io/driver/sqlite" // Sqlite driver based on CGO
	"gorm.io/gorm/logger"

	// "github.com/glebarez/sqlite" // Pure go SQLite driver, checkout https://github.com/glebarez/sqlite for details
	"gorm.io/gorm"
)

var db *gorm.DB

// Connect is a function to connect to the database
func Connect() {
	var err error

	dbFolder := strings.Split(config.ReadValue().DatabasePath, "/")
	err = os.Mkdir(strings.Join(dbFolder[:len(dbFolder)-1], "/"), 0755)
	if err != nil && !os.IsExist(err) {
		olog.Fatal(err.Error())
	}
	dbPath := config.ReadValue().DatabasePath
	db, err = gorm.Open(sqlite.Open(dbPath), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		olog.Error(err.Error())
		olog.Fatal("failed to connect database")
	}
	db.AutoMigrate(&entities.Password{})
}

// Get is a function to get the database connection
func Get() *gorm.DB {
	return db
}
