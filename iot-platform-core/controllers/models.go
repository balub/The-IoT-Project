package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lucsky/cuid"

	"github.com/balub/The-IoT-Project/databases"
	"github.com/balub/The-IoT-Project/databases/models"
)

type Field struct {
	Name     string `json:"name"`
	Type     string `json:"type"`
	Required string `json:"required"`
}

type ModelReqBody struct {
	ProjectID string  `json:"projectID"`
	Name      string  `json:"name"`
	Fields    []Field `json:"fields"`
}

func CreateModel(c *gin.Context) {
	// parseBody
	var body ModelReqBody

	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusPartialContent, gin.H{"message": "Error parsing body"})
		return
	}

	modelId := cuid.New()

	newModel := models.Models{ID: string(modelId), Name: body.Name, ProjectID: string(body.ProjectID)}

	errCase := databases.DB.Create(&newModel).Error

	if errCase != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": errCase.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "model created"})

}
