package setup

import (
	"golang-restful-api-framework/internal/config"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitSqlite(config config.Config) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(config.Database.Sqlite.DBPath), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return db
}
