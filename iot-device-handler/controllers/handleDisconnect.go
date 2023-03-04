package controllers

import (
	"fmt"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

var handleDisconnect mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
	fmt.Printf("Connect lost: %v", err)
}
