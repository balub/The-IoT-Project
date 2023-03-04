package databases

import (
	"github.com/balub/The-IoT-Project/databases/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (DB *gorm.DB)

func Connect(){
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
  if err != nil {
    panic("failed to connect database")
  }

	DB = db
}


func Migrate(){
	DB.AutoMigrate(&models.User{},&models.Projects{},&models.Devices{},&models.Models{},&models.Fields{})
}


