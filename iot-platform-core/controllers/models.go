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
	Required bool   `json:"required"`
}

type ModelReqBody struct {
	ProjectID string  `json:"projectID"`
	Name      string  `json:"name"`
	Fields    []Field `json:"fields"`
}

type dbModelInfo struct {
	ID string `json:"projectID"`
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

	for _, field := range body.Fields {
		newField := models.Fields{ModelId: string(modelId), Name: field.Name, Type: field.Type, Required: field.Required}
		errCase := databases.DB.Create(&newField).Error

		if errCase != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": errCase.Error()})
			return
		}
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "model created"})

}
