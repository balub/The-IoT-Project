package client

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
)

func FetchAll(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
	c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")
	if c.Request.Method == "OPTIONS" {
		c.AbortWithStatus(http.StatusOK)
		return
	}

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

	go func() {
		for {
			// Build the InfluxDB query
			startTime := time.Now().Add(-1 * time.Minute)
			stopTime := time.Now()
			query := fmt.Sprintf(`from(bucket:"%s")
				|> range(start: %d, stop: %d)
				|> filter(fn: (r) => r._measurement == "projectCore")
				|> pivot(rowKey: ["_time"], columnKey: ["_field"], valueColumn: "_value")
				`,
				influxBucket, startTime.UnixNano(), stopTime.UnixNano())

			// Execute the query and process the results
			result, err := queryAPI.Query(context.Background(), query)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to execute query"})
				return
			}

			// Parse the results and send them to the message channel
			for result.Next() {
				record := result.Record()
				fmt.Println(record)
				// messageChan <- record.ValueByKey("_value").(string)
			}

			// Sleep for one second before executing the query again
			time.Sleep(time.Second)
		}
	}()

	go func() {
		for {
			select {
			case <-c.Writer.CloseNotify():
				// The client has disconnected
				fmt.Println("disconnected from server")
				return
			case message := <-messageChan:
				fmt.Println(message)
				c.Writer.WriteString("data: " + message + "\n\n")
				c.Writer.Flush()
			}
		}
	}()
}
