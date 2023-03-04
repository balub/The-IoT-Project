package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email string `gorm:"unique" json:"email"`
	Password string `gorm:"size:256" json:"password"`
}

type Projects struct{
	gorm.Model
	Name string `gorm:"unique" json:"name"`
	DbUrl string `json:"dbUrl"`
	UserID int64 `json:"userID"`
	User User `gorm:"foreignKey:UserID"`
}

type Devices struct{
	gorm.Model
	Name string `json:"name"`
	AuthKey string `json:"authKey"`
	ProjectID int64 `json:"projectID"`
	Projects Projects `gorm:"foreignKey:ProjectID"`
}

type Models struct{
	gorm.Model
	Name	string `gorm:"uniqueIndex:idx_name_projectid" json:"name"`
	ProjectID int64 `gorm:"uniqueIndex:idx_name_projectid" json:"projectId"`
	Projects Projects `gorm:"foreignKey:ProjectID"`
}

type Fields struct{
	gorm.Model
	Name string `gorm:"uniqueIndex:idx_name_fieldname" json:"name"`
	Type string `json:"type"`
	Required bool `json:"required"`
	ModelId int64 `gorm:"uniqueIndex:idx_name_fieldname" json:"modelId"`
	Models Models `gorm:"foreignKey:ModelId"`
}