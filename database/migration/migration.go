package migration

import (
	"github.com/kaikeventura/golang-api/model"
	"gorm.io/gorm"
	"log"
)

func RunMigration(database *gorm.DB) {
	err := database.AutoMigrate(model.Book{})
	if err != nil {
		log.Fatal("Migration error: ", err)
	}
}
