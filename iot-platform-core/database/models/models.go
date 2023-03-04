package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email string `gorm:"unique"`
	Password string `gorm:"size:256"`
}

type Projects struct{
	gorm.Model
	Name string `gorm:"unique"`
	DbUrl string
	UserID int64
	User User `gorm:"foreignKey:UserID"`
}

type Devices struct{
	gorm.Model
	Name string
	AuthKey string
	ProjectID int64
	Projects Projects `gorm:"foreignKey:ProjectID"`
}

type Models struct{
	gorm.Model
	Name	string `gorm:"uniqueIndex:idx_name_projectid"`
	ProjectID int64 `gorm:"uniqueIndex:idx_name_projectid"`
	Projects Projects `gorm:"foreignKey:ProjectID"`
}

type Fields struct{
	gorm.Model
	Name string `gorm:"uniqueIndex:idx_name_fieldname"`
	Type string
	Required bool
	ModelId int64 `gorm:"uniqueIndex:idx_name_fieldname"`
	Models Models `gorm:"foreignKey:ModelId"`
}