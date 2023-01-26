package db

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/diogocarasco/golang-api-template-v2/config"

	_ "github.com/lib/pq"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var ConnPool = GetConnectionPool()

func initConnectionPool() (*sql.DB, error) {

	config := config.GetDB()

	db, err := gorm.Open(postgres.Open(config.GetConnectionString()), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	sqlDB, err := db.DB()
	if err != nil {
		fmt.Println("database error")
	} else {
		sqlDB.SetMaxIdleConns(10)
		sqlDB.SetMaxOpenConns(100)
		sqlDB.SetConnMaxLifetime(time.Hour)
	}

	if err = sqlDB.Ping(); err != nil {
		fmt.Println("database unreachable")
	}
	fmt.Println(sqlDB.Stats())
	return sqlDB, err

}

func GetConnectionPool() (db *sql.DB, err error) {
	db, err = initConnectionPool()
	if err != nil {
		return nil, err
	}

	return db, err
}
