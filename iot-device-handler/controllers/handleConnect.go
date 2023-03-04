package controllers

import (
	"fmt"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

var handleConnect mqtt.OnConnectHandler = func(client mqtt.Client) {
	fmt.Println("Connected")
}
