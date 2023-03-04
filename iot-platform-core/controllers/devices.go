package controllers

import (
	"fmt"
	"net/http"

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

type ProjectInfo struct {
	ProjectID string `json:"projectID"`
}

func CreateNewDevice(c *gin.Context) {
	// parseBody
	var body DeviceInput

	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusPartialContent, gin.H{"message": "Error parsing body"})
		return
	}

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

func FetchDevices(c *gin.Context) {
	var body ProjectInfo

	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusPartialContent, gin.H{"message": "Error parsing body"})
		return
	}

	// fetch user
	var devices []models.Devices
	databases.DB.First(&devices, fmt.Sprintf("ProjectID=%v", body.ProjectID))

	c.IndentedJSON(http.StatusOK, devices)

}
