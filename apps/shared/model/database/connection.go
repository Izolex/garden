package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func NewSQLite(path string) (*gorm.DB, *sql.DB) {
	return newConnection(sqlite.Open(path + "?parseTime=true"))
}

func NewMySQL(dsn string) (*gorm.DB, *sql.DB) {
	return newConnection(mysql.Open(dsn + "?parseTime=true"))
}

func newConnection(dialector gorm.Dialector) (*gorm.DB, *sql.DB) {
	db, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		panic(err)
	}
	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}

	return db, sqlDB
}
