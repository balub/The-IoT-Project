package client

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func SseHandler(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
	c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")
	c.Writer.Header().Set("Content-Type", "text/event-stream")
	c.Writer.Header().Set("Cache-Control", "no-cache")
	c.Writer.Header().Set("Connection", "keep-alive")

	messageChan := make(chan string, 1)

	go func() {
		for {
			fmt.Println("bro inside the routine")
			message := "macha yena da"
			messageChan <- message
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
