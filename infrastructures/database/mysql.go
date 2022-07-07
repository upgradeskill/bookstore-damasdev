package database

import (
	"database/sql"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func GetConnection() *sql.DB {

	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := "root:@tcp(127.0.0.1:3306)/golang?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
	})

	if err != nil {
		panic(err)
	}

	// Get generic database object sql.DB to use its functions
	sqlDB, err := db.DB()

	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(60 * time.Minute)

	// SetConnMaxIdleTime sets the maximum amount of time a connection may be idle.
	sqlDB.SetConnMaxIdleTime(5 * time.Minute)

	return sqlDB
}
