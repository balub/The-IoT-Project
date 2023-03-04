package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lucsky/cuid"

	"github.com/balub/The-IoT-Project/databases"
	"github.com/balub/The-IoT-Project/databases/models"
)

type ProjectInput struct {
	Name          string `json:"name"`
	DbUrl         string `json:"dbUrl"`
	DbAuthKey     string `json:"dbAuthKey"`
	BucketName    string `json:"bucketName"`
	DbProjectName string `json:"dbProjectName"`
}

func CreateNewProject(c *gin.Context) {
	// check for userID
	userID, exists := c.Get("userId")

	if !exists {
		c.JSON(http.StatusForbidden, gin.H{"message": "userId is required"})
		return
	}

	// fetch user
	var user models.User
	databases.DB.First(&user, userID)

	// parseBody
	var body ProjectInput

	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusPartialContent, gin.H{"message": "Error parsing body"})
		return
	}

	newProject := models.Projects{ID: cuid.New(), Name: body.Name, DbUrl: body.DbUrl, DbAuthKey: body.DbAuthKey, BucketName: body.BucketName, DbProjectName: body.DbProjectName, UserID: string(user.ID)}
	databases.DB.Create(&newProject)

	c.JSON(200, gin.H{"message": "project created"})
}

func FetchProjects(c *gin.Context) {
	userID, exists := c.Get("userId")
	fmt.Println(userID)

	if !exists {
		c.JSON(http.StatusForbidden, gin.H{"message": "userId is required"})
		return
	}

	// fetch user
	var user models.User
	databases.DB.First(&user, userID)

	var projects []models.Projects
	databases.DB.Where("user_id = ?", user.ID).Find(&projects)

	c.IndentedJSON(http.StatusOK, projects)
}
