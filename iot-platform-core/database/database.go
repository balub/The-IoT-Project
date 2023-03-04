package databases

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (DB *gorm.DB)

func Connect(){
	DB, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
  if err != nil {
    panic("failed to connect database")
  }
}


func Migrate(){
	DB.AutoMigrate()
}


