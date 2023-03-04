package controllers

// import (
// 	"fmt"
// 	"net/http"
// 	"strings"

// 	"github.com/gin-gonic/gin"
// 	"github.com/lucsky/cuid"

// 	"github.com/balub/The-IoT-Project/databases"
// 	"github.com/balub/The-IoT-Project/databases/models"
// 	"github.com/balub/The-IoT-Project/utils"
// )

// type DeviceInput struct {
// 	Name      string `json:"name"`
// 	AuthKey   string `json:"authKey"`
// 	ProjectID string `json:"projectID"`
// }

// func createNewDevice(c *gin.Context) {
// 	// check for userID
// 	userID, exists := utils.ExtractTokenID(c)

// 	if exists != nil {
// 		c.JSON(http.StatusForbidden, gin.H{"message": "userId is required"})
// 		return
// 	}

// 	userID = strings.ReplaceAll(userID, "\"", "'")

// 	// parseBody
// 	var body ProjectInput

// 	if err := c.BindJSON(&body); err != nil {
// 		c.JSON(http.StatusPartialContent, gin.H{"message": "Error parsing body"})
// 		return
// 	}

// 	// newDevice := models.Devices{ID: cuid.New(), }

// }
