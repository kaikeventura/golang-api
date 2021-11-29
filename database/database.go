package database

import (
	"github.com/kaikeventura/golang-api/database/migration"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"time"
)

var db *gorm.DB

func StartDatabase() {
	connection := "host=localhost port=25432 user=admin dbname=books password=123456"

	database, err := gorm.Open(postgres.Open(connection), &gorm.Config{})
	if err != nil {
		log.Fatal("Error: ", err)
	}

	db = database
	configuration, _ := db.DB()
	configuration.SetMaxIdleConns(10)
	configuration.SetMaxOpenConns(100)
	configuration.SetConnMaxLifetime(time.Hour)

	migration.RunMigration(db)
}

func CloseConnection() error {
	config, err := db.DB()
	if err != nil {
		return err
	}

	err = config.Close()
	if err != nil {
		return err
	}

	return nil
}

func GetDatabase() *gorm.DB {
	return db
}
