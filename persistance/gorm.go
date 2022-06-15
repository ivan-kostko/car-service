package persistance

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func DefaultSqliteDb(file string) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(file), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return db
}
