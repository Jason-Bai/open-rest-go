package models

import (
	"fmt"
	"log"
	"net/url"

	"github.com/Jason-Bai/open-rest-go/config"
	"github.com/Sirupsen/logrus"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var _db *gorm.DB

func openDB() *gorm.DB {
	dbUser := config.Config.DB.User
	dbPass := config.Config.DB.Pass
	dbHost := config.Config.DB.Host
	dbPort := config.Config.DB.Port
	dbName := config.Config.DB.Name

	connection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)

	params := url.Values{}
	params.Add("parseTime", "1")
	params.Add("loc", "Asia/Shanghai")

	dsn := fmt.Sprintf("%s?%s", connection, params.Encode())

	db, err := gorm.Open("mysql", dsn)

	if err != nil {
		log.Fatal("Database connection failed. Database error: %", err)
	}

	return db
}

// OpenDB open a connection to mysql database
func OpenDB(log *logrus.Logger) *gorm.DB {
	_db = openDB()

	// only for debug mode
	if "debug" == config.Config.MODE {
		autoMigration()
	}

	// set database logger
	_db.SetLogger(log)

	return _db
}

// AutoMigration auto migration custom tables
func autoMigration() {
	tables := []interface{}{
		&User{},
	}

	_db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(tables...)
}

// Close close connect to mysql database
func Close() {
	_db.Close()
}
