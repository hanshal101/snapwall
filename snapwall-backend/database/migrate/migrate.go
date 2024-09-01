package migrate

import (
	"log"

	"github.com/hanshal101/snapwall-backend/models"
	"gorm.io/gorm"
)

func MigrateModels(DB *gorm.DB) {
	if DB == nil {
		log.Fatalf("Error in migrating DB: DB == nil\n")
		return
	}
	DB.AutoMigrate(&models.Policy{})
	DB.AutoMigrate(&models.IP{})
	DB.AutoMigrate(&models.Port{})
	log.Println("DB Migrated Successfully")
}
