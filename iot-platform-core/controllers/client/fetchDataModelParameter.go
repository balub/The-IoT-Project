package client

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"

	"github.com/balub/The-IoT-Project/databases"
	"github.com/balub/The-IoT-Project/databases/models"
)

func FetchDataModelParameter(c *gin.Context) {

	projectToken := c.Param("projectToken")
	dataModel := c.Param("dataModel")
	parameter := c.Param("parameter")

	var existingProject models.Projects
	err := databases.DB.Find(&existingProject, fmt.Sprintf("id='%v'", projectToken)).Error

	if err != nil {
		c.String(http.StatusNoContent, "no such project found")
		return
	}

	var existingModel models.Models
	dbErr := databases.DB.Find(&existingModel, fmt.Sprintf("id='%v'", dataModel)).Error

	if dbErr != nil {
		c.String(http.StatusNoContent, "no such model found")
		return
	}

	fmt.Println(projectToken, dataModel)

	influxOrg := "theproject"
	influxURL := existingProject.DbUrl
	influxToken := existingProject.DbAuthKey
	influxBucket := existingProject.BucketName
	fmt.Println(influxOrg)
	if influxToken == "" || influxBucket == "" || influxURL == "" {
		c.String(http.StatusPartialContent, "error please set your influx data first")
		return
	}

	client := influxdb2.NewClient(influxURL, influxToken)
	queryAPI := client.QueryAPI(influxOrg)

	query := fmt.Sprintf(`from(bucket:"%s")
	|> range(start: 0)
	|> filter(fn: (r) => r["_measurement"] == "projectCore" and r["modelName"]=="%s" and r["dataModel"]=="%s")
	|> last()
	|> keep(columns:["%s"])
	`, influxBucket, projectToken, existingModel.Name, parameter)

	result, err := queryAPI.Query(context.Background(), query)
	if err != nil {
		c.String(http.StatusInternalServerError, "error Failed to execute query")
		fmt.Println("influx query issue", err)
		return
	}

	for result.Next() {
		// record := result.Record()
		values := result.Record().Values()

		resultJSON, err := json.Marshal(values)

		fmt.Println(string(resultJSON))

		if err != nil {
			// c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed parsing json"})
			fmt.Println("json parsing issue")
			return
		}

		c.String(200, string(resultJSON))
	}
}
