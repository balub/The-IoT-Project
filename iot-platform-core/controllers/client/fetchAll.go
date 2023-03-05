package client

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
)

func FetchAll(c *gin.Context) {
	c.Writer.Header().Set("Content-Type", "text/event-stream")
	c.Writer.Header().Set("Cache-Control", "no-cache")
	c.Writer.Header().Set("Connection", "keep-alive")

	messageChan := make(chan string, 1)

	// Create a new client and query API instance
	influxURL := "http://localhost:8086"
	influxToken := "IDkzBoF0Fmfgvu96O60rxU8SZy_Oz7wtABxlVlPWExQdAuHcALSrw4NMzSL3RVkxl3nh1NKwJQ92tC5vMghUQw=="
	influxOrg := "theproject"
	influxBucket := "nodemcu"
	client := influxdb2.NewClient(influxURL, influxToken)
	queryAPI := client.QueryAPI(influxOrg)

	closeNotify := c.Writer.(http.CloseNotifier).CloseNotify()
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	go func() {
		for {
			query := fmt.Sprintf(`from(bucket:"%s")
			|> range(start: 0)
			|> filter(fn: (r) => r["_measurement"] == "projectCore" and r["modelName"]=="sjr")
			|> last()
			`, influxBucket)

			// Execute the query and process the results
			result, err := queryAPI.Query(context.Background(), query)
			if err != nil {
				// c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to execute query"})
				fmt.Println("influx query issue")
				return
			}

			// Parse the results and send them to the message channel
			for result.Next() {
				// record := result.Record()
				values := result.Record().Values()

				resultJSON, err := json.Marshal(values)

				if err != nil {
					// c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed parsing json"})
					fmt.Println("json parsing issue")
					return
				}

				messageChan <- string(resultJSON)
			}

			// Sleep for one second before executing the query again
			time.Sleep(time.Second * 15)
		}
	}()

	go func() {
		for {
			select {
			case <-closeNotify:
				// The client has disconnected
				fmt.Println("disconnected from server")
				return
			case message := <-messageChan:
				fmt.Println(message, "sdgz")
				c.Writer.WriteString("data: " + message + "\n\n")
				c.Writer.Flush()
			}
		}
	}()
}
