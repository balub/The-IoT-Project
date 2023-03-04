package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID       string `gorm:"primarykey"`
	Email    string `gorm:"unique" json:"email"`
	Password string `gorm:"size:256;not null" json:"password"`
}

type Projects struct {
	gorm.Model
	ID            string `gorm:"primarykey;unique"`
	Name          string `gorm:"unique" json:"name"`
	DbUrl         string `gorm:"not null" json:"dbUrl"`
	DbAuthKey     string `gorm:"not null" json:"dbAuthKey"`
	BucketName    string `gorm:"not null" json:"bucketName"`
	DbProjectName string `gorm:"not null" json:"dbProjectName"`
	UserID        string `gorm:"not null" json:"userID"`
	User          User   `gorm:"foreignKey:UserID"`
}

type Devices struct {
	gorm.Model
	Name      string   `json:"name"`
	AuthKey   string   `json:"authKey"`
	ProjectID int64    `json:"projectID"`
	Projects  Projects `gorm:"foreignKey:ProjectID"`
}

type Models struct {
	gorm.Model
	Name      string   `gorm:"uniqueIndex:idx_name_projectid" json:"name"`
	ProjectID int64    `gorm:"uniqueIndex:idx_name_projectid" json:"projectId"`
	Projects  Projects `gorm:"foreignKey:ProjectID"`
}

type Fields struct {
	gorm.Model
	Name     string `gorm:"uniqueIndex:idx_name_fieldname" json:"name"`
	Type     string `json:"type"`
	Required bool   `json:"required"`
	ModelId  int64  `gorm:"uniqueIndex:idx_name_fieldname" json:"modelId"`
	Models   Models `gorm:"foreignKey:ModelId"`
}
