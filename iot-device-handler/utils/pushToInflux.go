package utils

import (
	"time"

	influxdb2 "github.com/influxdata/influxdb-client-go"
)

func PushToInflux(fields map[string]interface{}) error {

	// Set InfluxDB connection options
	options := influxdb2.DefaultOptions()
	options.SetBatchSize(20)

	// Create a new InfluxDB client
	client := influxdb2.NewClient("http://localhost:8086", "IDkzBoF0Fmfgvu96O60rxU8SZy_Oz7wtABxlVlPWExQdAuHcALSrw4NMzSL3RVkxl3nh1NKwJQ92tC5vMghUQw==")

	// Create a new WriteAPI using the client
	writeAPI := client.WriteAPI("theproject", "nodemcu")

	// Create a new Point object with measurement "cpu" and tags "host" and "region"
	point := influxdb2.NewPoint("projectCore", nil, fields, time.Now())

	// Write the data point to InfluxDB
	writeAPI.WritePoint(point)

	client.Close()

	return nil
}
