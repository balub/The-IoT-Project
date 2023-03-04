package controllers

import (
	"fmt"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

var HandleConnect mqtt.OnConnectHandler = func(client mqtt.Client) {
	fmt.Println("Connected")
}
