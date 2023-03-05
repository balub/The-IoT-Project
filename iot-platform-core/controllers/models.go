package controllers

import (
	"fmt"
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

type UpdateReqBody struct {
	ModelID string  `json:"modelID"`
	Fields  []Field `json:"fields"`
}

type dbModelInfo struct {
	ID string `json:"projectID"`
}

type FetchListModel struct {
	ID   string `json:"modelID"`
	Name string `json:"name"`
}

type FetchListField struct {
	Name     string `json:"name"`
	Type     string `json:"type"`
	Required bool   `json:"required"`
}

type FinalDataModel struct {
	ID     string           `json:"modelID"`
	Name   string           `json:"name"`
	Fields []FetchListField `json:"fields"`
}

type UpdateDataModelBody struct {
	ModelID string  `json:"name"`
	Fields  []Field `json:"fields"`
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

func FetchModelInfo(c *gin.Context) {
	var body ModelReqBody

	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusPartialContent, gin.H{"message": "Error parsing body"})
		return
	}

	var models []FetchListModel
	databases.DB.Table("models").Select("id, name").Where(fmt.Sprintf("project_id='%v'", body.ProjectID)).Scan(&models)

	var arrLength = len(models)

	finalData := make([]FinalDataModel, arrLength)

	for index, model := range models {
		finalData[index].ID = model.ID
		finalData[index].Name = model.Name
	}

	for _, model := range models {

		var modelID = model.ID
		var fields []FetchListField
		databases.DB.Table("fields").Select("name, type, required").Where(fmt.Sprintf("model_id='%v'", modelID)).Scan(&fields)

		for index, field := range fields {
			finalData[index].Fields = append(finalData[index].Fields, FetchListField{field.Name, field.Type, field.Required})
		}
	}

	c.IndentedJSON(http.StatusOK, finalData)
}

func UpdateDataModel(c *gin.Context) {
	var body UpdateReqBody

	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusPartialContent, gin.H{"message": "Error parsing body"})
		return
	}

	var field models.Fields
	deleteError := databases.DB.Where(fmt.Sprintf("model_id='%v'", body.ModelID)).Delete(&field).Error

	if deleteError != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Delete error": deleteError.Error()})
		return
	}

	for _, field := range body.Fields {
		newField := models.Fields{ModelId: string(body.ModelID), Name: field.Name, Type: field.Type, Required: field.Required}
		errCase := databases.DB.Create(&newField).Error

		if errCase != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Insert error": errCase.Error()})
			return
		}
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "model updated"})
}
