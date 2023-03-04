// package main

// import (
// 	"context"
// 	"fmt"
// 	"log"

// 	influxdb2 "github.com/influxdata/influxdb-client-go"
// )

// func main() {
// 	// Set InfluxDB connection options
// 	options := influxdb2.DefaultOptions()
// 	options.SetBatchSize(20)

// 	// Create a new InfluxDB client
// 	client := influxdb2.NewClient("http://localhost:8086", "a2x2Ql81OyjHep3TxbyCTE-wGumm2EjMWwML5abM9Y2jAee4UbdBGuo6lkOuKBhMDrppi4Z7Wyr432SE99PXbw==")

// 	// Create a new WriteAPI using the client
// 	writeAPI := client.WriteAPI("theproject", "nodemcu")

// 	// Create a new Point object with measurement "cpu" and tags "host" and "region"
// 	// tags := map[string]string{"host": "server01", "region": "us-west"}
// 	fields := map[string]interface{}{"light":true}
// 	point := influxdb2.NewPoint("cpu", nil, fields, time.Now())

// 	// Write the data point to InfluxDB
// 	writeAPI.WritePoint(point)

//     // query

// 	// queryAPI := client.QueryAPI("theproject")

// 	// // Construct a Flux query to select all data points from the "cpu" measurement
// 	// fluxQuery := fmt.Sprintf(`from(bucket:"nodemcu")
//     //                          |> range(start: -1h)
//     //                          |> filter(fn: (r) => r._measurement == "cpu")`)

// 	// result, err := queryAPI.Query(context.Background(), fluxQuery)
// 	// if err != nil {
// 	// 	log.Fatal(err)
// 	// }

// 	// fmt.Println(result.TableMetadata())
// 	// for result.Next() {
// 	// 	// Get the data point as a map
// 	// 	point := result.Record().Values()

// 	// 	// Print the data point
// 	// 	fmt.Printf("Time: %v, Value: %v\n", point["_time"], point["light"])
// 	}
// 	// Close the InfluxDB client
// 	client.Close()
// }

package main

import (
	"github.com/balub/The-IoT-Project/mqtt"
)

func main() {

	mqtt.Subscriber()
}
