package controllers

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/lucsky/cuid"

	"github.com/balub/The-IoT-Project/databases"
	"github.com/balub/The-IoT-Project/databases/models"
	"github.com/balub/The-IoT-Project/utils"
)

type DeviceInput struct {
	Name      string `json:"name"`
	ProjectID string `json:"projectID"`
}

func CreateNewDevice(c *gin.Context) {
	// check for userID
	userID, exists := utils.ExtractTokenID(c)

	if exists != nil {
		c.JSON(http.StatusForbidden, gin.H{"message": "userId is required"})
		return
	}

	userID = strings.ReplaceAll(userID, "\"", "'")

	// parseBody
	var body DeviceInput

	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusPartialContent, gin.H{"message": "Error parsing body"})
		return
	}

	// fetch user
	// var project models.Projects
	// fetchErrorCase := databases.DB.First(&project, fmt.Sprintf("id=%v", body.ProjectID)).Error

	// if fetchErrorCase != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": "Project Id Not Found"})
	// 	return
	// }

	deviceId := cuid.New()

	token, tokenErr := utils.GenerateProjectToken(body.ProjectID, deviceId)

	if tokenErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": tokenErr.Error()})
		return
	}

	newDevice := models.Devices{ID: deviceId, Name: body.Name, AuthKey: token, ProjectID: string(body.ProjectID)}

	errCase := databases.DB.Create(&newDevice).Error

	if errCase != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": errCase.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "device created"})
}
